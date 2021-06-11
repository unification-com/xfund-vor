package chainlisten

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"math/rand"
	"oracle/config"
	"oracle/contracts/vor_coordinator"
	"oracle/models/database"
	"oracle/service"
	"oracle/tools/vor"
	"strings"
	"sync"
	"time"
)

type VORCoordinatorListener struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_coordinator.VorCoordinator
	query           ethereum.FilterQuery
	wg              *sync.WaitGroup
	service         *service.Service
	keyHash         [32]byte
	context         context.Context
	logger          *logrus.Logger
}

func NewVORCoordinatorListener(contractHexAddress string, ethHostAddress string,
	service *service.Service, logger *logrus.Logger, ctx context.Context) (*VORCoordinatorListener, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(contractHexAddress)
	instance, err := vor_coordinator.NewVorCoordinator(contractAddress, client)
	if err != nil {
		return nil, err
	}

	var lastBlock *big.Int
	lastRequest, err := service.Store.Db.GetLast()
	if blockNumber, _ := service.Store.Keystorage.GetBlockNumber(); blockNumber != 0 {
		lastBlock = big.NewInt(blockNumber + 1)
	} else if lastRequest.GetRequestBlockNumber() != 0 {
		lastBlock = big.NewInt(int64(lastRequest.GetRequestBlockNumber()))
	} else if config.Conf.FirstBlockNumber != 0 {
		lastBlock = big.NewInt(int64(config.Conf.FirstBlockNumber + 1))
	} else {
		lastBlock = big.NewInt(1)
	}

	keyHash, err := service.VORCoordinatorCaller.HashOfKey()
	return &VORCoordinatorListener{
		client:          client,
		contractAddress: contractAddress,
		instance:        instance,
		query: ethereum.FilterQuery{
			FromBlock: lastBlock,
			Addresses: []common.Address{contractAddress},
		},
		service: service,
		context: ctx,
		keyHash: keyHash,
		wg:      &sync.WaitGroup{},
		logger:  logger,
	}, err
}

func (d VORCoordinatorListener) StartPoll() (err error) {
	d.wg.Add(1)

	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "StartPoll",
		"action":     "begin polling",
		"from_block": d.query.FromBlock.Uint64(),
	}).Info()

	var sleepTime = int32(30)
	if config.Conf.CheckDuration != 0 {
		sleepTime = config.Conf.CheckDuration
	}
	for {
		err = d.ProcessIncommingEvents()
		err = d.CheckStuck()
		time.Sleep(time.Duration(rand.Int31n(sleepTime)) * time.Second)
	}
	d.wg.Wait()
	return
}

func (d VORCoordinatorListener) Shutdown() {
	d.wg.Done()
}

func (d *VORCoordinatorListener) SetLastBlockNumber(blockNumber uint64) (err error) {
	d.query.FromBlock = big.NewInt(int64(blockNumber + 1))
	err = d.service.Store.Keystorage.SetBlockNumber(int64(blockNumber))
	return
}

func (d *VORCoordinatorListener) processStuckRequest(request database.RandomnessRequest, currentBlockNum uint64) {

	blockDiff := currentBlockNum - request.GetRequestBlockNumber()
	fulfilTxHash := common.HexToHash(request.GetFulfillTxHash())

	d.logger.WithFields(logrus.Fields{
		"package":  "chainlisten",
		"function": "processStuckRequest",
		"action":   "log info",
		"request_id":  request.GetRequestId(),
		"tx_hash":  request.GetFulfillTxHash(),
		"req_block": request.GetRequestBlockNumber(),
		"curr_block": currentBlockNum,
		"block_diff": blockDiff,
	}).Info()

	// same block - ignore for now
	if blockDiff == 0 {
		return
	}

	// check if the block hash for the request has been stored
	foundHash, _, bhErr := d.service.VORCoordinatorCaller.GetBlockHashFromBlockStore(request.GetRequestBlockNumber())

	// used later to store failed fulfill tx history
	failedGasUsed := uint64(0)
	failedGasPrice := uint64(0)

	// only relevant if the tx was broadcast
	if len(request.GetFulfillTxHash()) > 0 {
		// check if it's pending
		lastFulfillTx, isPending, err := d.client.TransactionByHash(context.Background(), fulfilTxHash)

		if err != nil {
			// probably not in Tx pool yet
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "get fulfill tx",
				"request_id":  request.GetRequestId(),
				"tx_hash":  request.GetFulfillTxHash(),
			}).Error(err.Error())
			return
		}

		// no point continuing if it's still pending. Log it and move on.
		// todo - However, if block diff is high, store block hash in block hash store contract
		//        and resend using same tx nonce, with higher gas price
		if isPending {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "get fulfill tx",
				"request_id":  request.GetRequestId(),
				"tx_hash":  request.GetFulfillTxHash(),
			}).Info("tx still pending - ignore")
			return
		}

		// grab the actual gas price used to broadcast the tx
		failedGasPrice = lastFulfillTx.GasPrice().Uint64()
	}

	// is the request really old?
	if blockDiff > 250 {
		// check if hash is in block store
		if !foundHash || bhErr != nil {
			// block hash not in store. Flag as failed and move on
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "check request age",
				"request_id":  request.GetRequestId(),
			}).Info("request too old and block hash not in block store")
			_ = d.service.Store.Db.UpdateRequestStatus(request.GetRequestId(), database.REQUEST_STATUS_FULFILMENT_FAILED, "request too old and block hash not in block store")
			return
		}
	}

	// get original fail reason stored in db. This may be a tx that was not broadcast due to
	// out of gas, nonce too low etc.
	failReason := request.GetStatusReason()

	// if it was broadcast, however, try to get the receipt and determine the fail reason
	if request.GetStatus() == database.REQUEST_STATUS_SENT {
		// try and get the receipt
		fulfillReceipt, err := d.client.TransactionReceipt(context.Background(), fulfilTxHash)
		if err != nil {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "get fulfil tx receipt",
				"request_id":  request.GetRequestId(),
				"tx_hash":  request.GetFulfillTxHash(),
			}).Error(err.Error())
			return
		}

		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "processStuckRequest",
			"action":   "got fulfil tx receipt",
			"request_id":  request.GetRequestId(),
			"tx_hash":  request.GetFulfillTxHash(),
			"status": fulfillReceipt.Status,
		}).Info()

		if fulfillReceipt.Status == 1 {
			// oddly, seems OK, but this should probably never occur.
			// todo - check why event wasn't recorded
			return
		}

		// generic reason for now
		// todo - try to get actual revert reason
		failReason = "transaction reverted"
		failedGasUsed = fulfillReceipt.GasUsed
	}

	// Add fail info to failed Tx history table
	_ = d.service.Store.Db.InsertNewFailedFulfilment(request.GetRequestId(), request.GetFulfillTxHash(), failedGasUsed, failedGasPrice, failReason)

	// at some point, we just have to stop trying...
	if request.GetFulfillmentAttempts() >= 3 {
		// too many fails
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "processStuckRequest",
			"action":   "check num attempts",
			"request_id":  request.GetRequestId(),
			"num_attempts": request.GetFulfillmentAttempts(),
		}).Info()

		_ = d.service.Store.Db.UpdateRequestStatus(request.GetRequestId(), database.REQUEST_STATUS_FULFILMENT_FAILED, "too many failed attempts")
		return
	}

	// to be safe, if the request is getting old, store the block hash of the request Tx's block in block store contract
	// this ensures the request can still hopefully be fulfilled if more than 256 blocks pass during subsequent retries.
	if blockDiff > 50 && !foundHash {
		bhTx, err := d.service.VORCoordinatorCaller.StoreBlockHash(request.GetRequestBlockNumber())
		if err != nil {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "store blockhash in block store",
			}).Error(err.Error())
		} else {
			// add block store tx data to db table
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "processStuckRequest",
				"action":   "store blockhash in block store",
				"block_num": request.GetRequestBlockNumber(),
				"block_hash": request.GetRequestBlockHash(),
			}).Info("store tx sent")
			_ = d.service.Store.Db.InsertNewStoredBlock(request.GetRequestBlockHash(), request.GetRequestBlockNumber(), bhTx.Hash().Hex())
		}
	}

	// retry sending fulfilment Tx and update DB as appropriate.
	// We'll use the original seed, block hash data etc. but the proof
	// generated will be unique
	seedBig, _ := hexutil.DecodeBig(request.Seed)
	byteSeed, err := vor.BigToSeed(seedBig)
	reqBlockHash := common.HexToHash(request.GetRequestBlockHash())

	// retry fulfillment
	fTx, err := d.service.FulfillRandomness(byteSeed, reqBlockHash, int64(request.GetRequestBlockNumber()))

	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processStuckRequest",
			"action":     "retry fulfill request",
			"request_id": request.GetRequestId(),
		}).Error(err.Error())
		// possibly failed due to gas too low, or nonce too low. Flag so we can try again later
		_ = d.service.Store.Db.UpdateRequestStatus(request.GetRequestId(), database.REQUEST_STATUS_TX_FAILED, err.Error())
	} else {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processStuckRequest",
			"action":     "retry fulfill request",
			"request_id": request.GetRequestId(),
			"tx_hash":    fTx.Hash().Hex(),
		}).Info("fulfill tx sent")
		// update the Db - Tx successfully broadcast.
		_ = d.service.Store.Db.UpdateFulfilmentSent(request.GetRequestId(), database.REQUEST_STATUS_SENT, fTx.Hash().Hex())
	}

	return
}

func (d *VORCoordinatorListener) CheckStuck() error {

	d.logger.WithFields(logrus.Fields{
		"package":  "chainlisten",
		"function": "CheckStuck",
		"action":   "check stuck fulfilments",
	}).Info()

	currentBlockNum, err := d.client.BlockNumber(context.Background())

	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "CheckStuck",
			"action":   "get block num",
		}).Error(err.Error())
		return err
	}

	// get requests status = SENT || FAILED_TX from request_randomness table
	requests, err := d.service.Store.Db.GetStuckOrFailedTx()

	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "CheckStuck",
			"action":   "get stuck or failed tx requests",
		}).Error(err.Error())
		return err
	}

	for _, request := range requests {
		// process
		d.processStuckRequest(request, currentBlockNum)
	}

	return nil
}

func (d *VORCoordinatorListener) ProcessIncommingEvents() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	if len(logs) == 0 {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "ProcessIncommingEvents",
			"action":     "check events",
			"from_block": d.query.FromBlock.Uint64(),
		}).Info("no applicable logs")

		thisBlockNum, err := d.client.BlockNumber(context.Background())
		if err == nil {
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "ProcessIncommingEvents",
				"action":     "set last block",
				"from_block": thisBlockNum - 1,
			}).Info("no applicable logs")
			_ = d.SetLastBlockNumber(thisBlockNum - 1)
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "ProcessIncommingEvents",
				"action":   "get block num",
			}).Error(err.Error())
		}
		return nil
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(vor_coordinator.VorCoordinatorABI)))
	if err != nil {
		return err
	}

	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	logRandomnessRequestFulfilledHash := crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))

	for index, vLog := range logs {

		d.logger.WithFields(logrus.Fields{
			"package":   "chainlisten",
			"function":  "ProcessIncommingEvents",
			"action":    "log",
			"block_num": vLog.BlockNumber,
			"log_index": vLog.Index,
		}).Info()

		gasPrice := uint64(0)
		gasUsed := uint64(0)

		txRec, err := d.client.TransactionReceipt(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need to clean up and gather any missing data if Tx query above fails
			gasUsed = txRec.GasUsed
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "ProcessIncommingEvents",
				"action":   "get TransactionReceipt",
			}).Error(err.Error())
		}

		tx, _, err := d.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need to clean up and gather any missing data if Tx query above fails
			gasPrice = tx.GasPrice().Uint64()
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "ProcessIncommingEvents",
				"action":   "get TransactionByHash",
			}).Error(err.Error())
		}

		if index == len(logs)-1 {
			err = d.SetLastBlockNumber(vLog.BlockNumber)
		}

		switch vLog.Topics[0].Hex() {
		// RandomnessRequest event detected
		case logRandomnessRequestHash.Hex():
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "ProcessIncommingEvents",
				"action":     "check event name",
				"event_name": "RandomnessRequest",
			}).Info("processing event")

			event := vor_coordinator.VorCoordinatorRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "ProcessIncommingEvents",
					"action":   "UnpackIntoInterface",
				}).Error(err.Error())
				return err
			}

			if event.KeyHash == d.keyHash {

				requestId := common.Bytes2Hex(event.RequestID[:])

				d.logger.WithFields(logrus.Fields{
					"package":    "chainlisten",
					"function":   "ProcessIncommingEvents",
					"action":     "check event keyhash",
					"request_id": requestId,
				}).Info("It's a request for me =)")

				byteSeed, err := vor.BigToSeed(event.Seed)

				if err != nil {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "ProcessIncommingEvents",
						"action":   "BigToSeed",
					}).Error(err.Error())
					return err
				}

				// check status and if requests already exists
				reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

				if reqDbRes.ID == 0 {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "ProcessIncommingEvents",
						"action":   "check db for request",
					}).Info("new request")

					seedHex := hexutil.EncodeBig(event.Seed)

					err = d.service.Store.Db.InsertNewRequest(
						common.Bytes2Hex(event.KeyHash[:]),
						seedHex, event.Sender.Hex(),
						requestId,
						database.REQUEST_STATUS_INITIALISED,
						vLog.BlockHash.Hex(),
						vLog.BlockNumber,
						vLog.TxHash.Hex(),
						gasUsed,
						gasPrice,
						event.Fee.Uint64(),
					)

					fulfillTx, err := d.service.FulfillRandomness(byteSeed, vLog.BlockHash, int64(vLog.BlockNumber))

					if err != nil {
						d.logger.WithFields(logrus.Fields{
							"package":    "chainlisten",
							"function":   "ProcessIncommingEvents",
							"action":     "fulfill request",
							"request_id": requestId,
						}).Error(err.Error())
						_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_TX_FAILED, err.Error())
					} else {
						d.logger.WithFields(logrus.Fields{
							"package":    "chainlisten",
							"function":   "ProcessIncommingEvents",
							"action":     "fulfill request",
							"request_id": requestId,
							"tx_hash":    fulfillTx.Hash().Hex(),
						}).Info("fulfill tx sent")
						_ = d.service.Store.Db.UpdateFulfilmentSent(requestId, database.REQUEST_STATUS_SENT, fulfillTx.Hash().Hex())
					}
				} else {
					d.logger.WithFields(logrus.Fields{
						"package":    "chainlisten",
						"function":   "ProcessIncommingEvents",
						"action":     "check db for request",
						"request_id": reqDbRes.RequestId,
						"status":     reqDbRes.GetStatusString(),
					}).Info("request already in db")
				}
			} else {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "ProcessIncommingEvents",
					"action":   "check event keyhash",
				}).Info("Looks like it's not addressed to me =(")
			}
			continue
		// RandomnessRequestFulfilled event detected
		case logRandomnessRequestFulfilledHash.Hex():
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "ProcessIncommingEvents",
				"action":     "check event name",
				"event_name": "RandomnessRequestFulfilled",
			}).Info("processing event")

			event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data)
			requestId := common.Bytes2Hex(event.RequestId[:])
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "ProcessIncommingEvents",
				"action":     "check request exists",
				"request_id": requestId,
			}).Info()

			reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

			if reqDbRes.ID != 0 {
				d.logger.WithFields(logrus.Fields{
					"package":    "chainlisten",
					"function":   "ProcessIncommingEvents",
					"action":     "confirm fulfillment",
					"request_id": requestId,
				}).Info("confirmed request fulfilment for request")

				if err != nil {
					return err
				}

				err = d.service.Store.Db.UpdateFulfillment(
					requestId,
					database.REQUEST_STATUS_SUCCESS,
					event.Output.String(),
					vLog.BlockHash.Hex(),
					vLog.BlockNumber,
					vLog.TxHash.Hex(),
					gasUsed,
					gasPrice,
				)
				if err != nil {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "ProcessIncommingEvents",
						"action":   "UpdateFulfillment",
					}).Error(err.Error())
					return err
				}
			} else {
				d.logger.WithFields(logrus.Fields{
					"package":    "chainlisten",
					"function":   "ProcessIncommingEvents",
					"action":     "confirm fulfillment",
					"request_id": requestId,
				}).Warning("request id does not exist in db. Probably not mine")
			}
			continue
		default:
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "ProcessIncommingEvents",
				"action":   "check event name",
			}).Info("event not applicable")
			continue
		}
	}

	return err
}

func (d VORCoordinatorListener) RandomnessRequest() {

}
