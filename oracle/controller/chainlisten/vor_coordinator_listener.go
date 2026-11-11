package chainlisten

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

	lastBlock, err := startBlock(service)

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

func (d *VORCoordinatorListener) StartPoll() (err error) {
	d.wg.Add(1)

	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "StartPoll",
		"action":     "begin polling",
		"from_block": d.query.FromBlock.Uint64(),
	}).Info()

	// Poll cadence: a base interval plus bounded jitter, so the floor is ALWAYS non-zero. The old code slept
	// time.Sleep(rand.Int31n(base)) which returns [0,base) - frequently 0 - so a stuck job (e.g. a permanently
	// reverting under-fee request) re-ran the loop with no delay, pegging a core and churning the heap (the
	// 2026-06 CPU/RAM runaway). It also discarded every iteration error and never honoured ctx cancellation.
	base := int32(30)
	if config.Conf.CheckDuration > 0 {
		base = config.Conf.CheckDuration
	}
	for {
		if d.context.Err() != nil {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "StartPoll",
				"action":   "context cancelled - stopping poller",
			}).Info()
			return d.context.Err()
		}

		errEvents := d.ProcessIncommingEvents()
		errJobs := d.CheckJobs()

		// Always wait at least `base` seconds; back off further on an RPC/processing error so a struggling
		// or dead endpoint is not hammered in a tight loop.
		wait := base + rand.Int31n(base)
		if errEvents != nil || errJobs != nil {
			fields := logrus.Fields{"package": "chainlisten", "function": "StartPoll", "action": "poll iteration error - backing off"}
			if errEvents != nil {
				fields["err_events"] = errEvents.Error()
			}
			if errJobs != nil {
				fields["err_jobs"] = errJobs.Error()
			}
			d.logger.WithFields(fields).Warn()
			wait = base * 2
		}

		select {
		case <-d.context.Done():
			return d.context.Err()
		case <-time.After(time.Duration(wait) * time.Second):
		}
	}
}

func (d VORCoordinatorListener) Shutdown() {
	d.wg.Done()
}

// startBlock is where polling resumes: the block saved in the keystore, else the last request's block, else
// the configured first block, else 1. A saved block that isn't positive is ignored: older versions could save
// -1 (see SetLastBlockNumber).
func startBlock(service *service.Service) (*big.Int, error) {
	lastRequest, err := service.Store.Db.GetLast()
	if blockNumber, _ := service.Store.Keystorage.GetBlockNumber(); blockNumber > 0 {
		return big.NewInt(blockNumber), err
	}
	if lastRequest.GetRequestBlockNumber() != 0 {
		return big.NewInt(int64(lastRequest.GetRequestBlockNumber())), err
	}
	if config.Conf.FirstBlockNumber != 0 {
		return big.NewInt(int64(config.Conf.FirstBlockNumber)), err
	}
	return big.NewInt(1), err
}

// SetLastBlockNumber makes the next poll start from the block before blockNumber. A zero block number (e.g.
// from a misbehaving RPC) is ignored: blockNumber - 1 would wrap round to the maximum uint64, which used to be
// saved in the keystore as -1.
func (d *VORCoordinatorListener) SetLastBlockNumber(blockNumber uint64) (err error) {
	if blockNumber == 0 {
		return nil
	}
	d.query.FromBlock = big.NewInt(int64(blockNumber - 1))
	err = d.service.Store.Keystorage.SetBlockNumber(int64(blockNumber - 1))
	return
}

func (d *VORCoordinatorListener) recordBlockHash(blockNumber uint64, blockHash string) {
	// to be safe, if the request is getting old, store the block hash of the request Tx's block in block store contract
	// this ensures the request can still hopefully be fulfilled if more than 256 blocks pass during subsequent retries.
	bhTx, err := d.service.VORCoordinatorCaller.StoreBlockHash(blockNumber)
	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "recordBlockHash",
			"action":   "store blockhash in block store",
		}).Error(err.Error())
	} else {
		// add block store tx data to db table
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "recordBlockHash",
			"action":     "store blockhash in block store",
			"block_num":  blockNumber,
			"block_hash": blockHash,
		}).Info("store tx sent")
		_ = d.service.Store.Db.InsertNewStoredBlock(blockHash, blockNumber, bhTx.Hash().Hex())
	}
}

// isPermanentFulfilmentError reports whether a fulfilment send/estimate error is PERMANENT - one that will
// never succeed by retrying, so the request must be failed terminally rather than re-sent forever. The classic
// case is a request whose paid fee is below the oracle's registered fee: the coordinator reverts ("Below agreed
// payment"), which - with gas estimation enabled (GasLimit 0) - surfaces as an "execution reverted" error BEFORE
// broadcast. Genuinely transient errors (nonce too low, RPC blips) are NOT matched and fall through to the normal
// capped retry. This guard is what stops the fee-revert CPU/RAM runaway.
func isPermanentFulfilmentError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	for _, marker := range []string{"execution reverted", "below agreed payment", "always failing transaction"} {
		if strings.Contains(msg, marker) {
			return true
		}
	}
	return false
}

func (d *VORCoordinatorListener) processFulfillment(requestId string, requestTxReceipt *types.Receipt, currentBlockNum uint64) {
	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "processFulfillment",
		"action":     "begin request fulfillment",
		"request_id": requestId,
	}).Info()

	contractAbi, _ := abi.JSON(strings.NewReader(vor_coordinator.VorCoordinatorABI))
	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))

	for _, vLog := range requestTxReceipt.Logs {
		if vLog.Topics[0].Hex() == logRandomnessRequestHash.Hex() {
			event := vor_coordinator.VorCoordinatorRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				d.logger.WithFields(logrus.Fields{
					"package":    "chainlisten",
					"function":   "processFulfillment",
					"action":     "unpack RandomnessRequest event to abi",
					"request_id": requestId,
				}).Error(err.Error())
				return
			}
			byteSeed, err := vor.BigToSeed(event.Seed)
			seedHex := hexutil.EncodeBig(event.Seed)
			requestBlockHash := requestTxReceipt.BlockHash

			_ = d.service.Store.Db.UpdateRequestBlockAndSeed(requestId, requestBlockHash.Hex(), seedHex, requestTxReceipt.BlockNumber.Uint64())

			// send fulfillment
			fTx, err := d.service.FulfillRandomness(byteSeed, requestBlockHash, requestTxReceipt.BlockNumber.Uint64())
			if err != nil {
				// A permanent revert (e.g. the request's paid fee is below the registered fee) is caught
				// here by gas estimation BEFORE broadcast - fail it terminally so it drops out of the job
				// queue, instead of re-sending it forever (the fee-revert CPU/RAM runaway). Transient errors
				// (nonce/RPC) fall through to TX_FAILED and the normal capped retry.
				permanent := isPermanentFulfilmentError(err)
				status := database.REQUEST_STATUS_TX_FAILED
				if permanent {
					status = database.REQUEST_STATUS_FULFILMENT_FAILED
				}
				d.logger.WithFields(logrus.Fields{
					"package":    "chainlisten",
					"function":   "processFulfillment",
					"action":     "fulfill request",
					"request_id": requestId,
					"permanent":  permanent,
				}).Error(err.Error())
				_ = d.service.Store.Db.UpdateRequestStatus(requestId, status, err.Error())
			} else {
				d.logger.WithFields(logrus.Fields{
					"package":            "chainlisten",
					"function":           "processFulfillment",
					"action":             "fulfill request",
					"request_id":         requestId,
					"request_tx_hash":    requestTxReceipt.TxHash,
					"fulfill_tx_hash":    fTx.Hash().Hex(),
					"request_block_hash": requestBlockHash,
					"request_block_num":  requestTxReceipt.BlockNumber,
					"seed":               event.Seed,
				}).Info("fulfill tx sent")
				// update the Db - Tx successfully broadcast.
				_ = d.service.Store.Db.UpdateFulfilmentSent(requestId, database.REQUEST_STATUS_SENT, fTx.Hash().Hex(), currentBlockNum)
			}
		}
	}
	return
}

func (d *VORCoordinatorListener) processFailed(request database.RandomnessRequest, requestTxReceipt *types.Receipt, currentBlockNum uint64) {
	requestId := request.GetRequestId()
	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "processFailed",
		"action":     "start",
		"request_id": requestId,
	}).Info()

	// used later to store failed fulfill tx history
	failedGasUsed := request.GetFulfillGasUsed()
	failedGasPrice := request.GetFulfillGasPrice()
	failReason := request.GetStatusReason()

	// Add fail info to failed Tx history table
	_ = d.service.Store.Db.InsertNewFailedFulfilment(requestId, request.GetFulfillTxHash(), failedGasUsed, failedGasPrice, failReason)

	// at some point, we just have to stop trying...
	if request.GetFulfillmentAttempts() >= 3 {
		// too many fails
		d.logger.WithFields(logrus.Fields{
			"package":      "chainlisten",
			"function":     "processFailed",
			"action":       "check num attempts",
			"request_id":   requestId,
			"num_attempts": request.GetFulfillmentAttempts(),
		}).Info()

		_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, "too many failed attempts")
		return
	}

	requestBlockDiff := currentBlockNum - requestTxReceipt.BlockNumber.Uint64()

	// finally, try to send a new fulfillment
	// check if the block hash for the request has been stored in BlockHash contract
	foundHash, _, bhErr := d.service.VORCoordinatorCaller.GetBlockHashFromBlockStore(requestTxReceipt.BlockNumber.Uint64())

	// is the request really old?
	if requestBlockDiff > 250 {
		// check if hash is in block store
		if !foundHash || bhErr != nil {
			// block hash not in store. Flag as failed and move on
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "processFailed",
				"action":     "check request age",
				"request_id": request.GetRequestId(),
			}).Info("request too old and block hash not in block store")
			_ = d.service.Store.Db.UpdateRequestStatus(request.GetRequestId(), database.REQUEST_STATUS_FULFILMENT_FAILED, "request too old and block hash not in block store")
			return
		}
	}

	// record the block hash in the contract to be safe
	if requestBlockDiff > 50 && !foundHash {
		d.recordBlockHash(requestTxReceipt.BlockNumber.Uint64(), requestTxReceipt.BlockHash.Hex())
	}

	d.processFulfillment(requestId, requestTxReceipt, currentBlockNum)

	return
}

func (d *VORCoordinatorListener) processPossiblyStuck(request database.RandomnessRequest, requestTxReceipt *types.Receipt, currentBlockNum uint64) {
	requestId := request.GetRequestId()
	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "processPossiblyStuck",
		"action":     "start",
		"request_id": requestId,
	}).Info()

	requestBlockDiff := currentBlockNum - requestTxReceipt.BlockNumber.Uint64()
	lastFulfillSentBlockDiff := currentBlockNum - request.GetLastFulfillSentBlockNumber()
	if lastFulfillSentBlockDiff < 2 {
		// too soon - may take a while for Tx to be broadcast/picked up
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processPossiblyStuck",
			"action":     "check block diff since fulfill tx sent",
			"request_id": requestId,
			"block_diff": lastFulfillSentBlockDiff,
		}).Info("not enough blocks since last sent. Wait.")
		return
	}

	fulfilTxHash := common.HexToHash(request.GetFulfillTxHash())
	// check if it's pending
	lastFulfillTx, isPending, err := d.client.TransactionByHash(context.Background(), fulfilTxHash)

	if err != nil {
		// probably not in Tx pool yet
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processPossiblyStuck",
			"action":     "get fulfill tx",
			"request_id": requestId,
			"tx_hash":    request.GetFulfillTxHash(),
		}).Error(err.Error())
		return
	}

	// no point continuing if it's still pending. Log it and move on.
	if isPending {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processPossiblyStuck",
			"action":     "check fulfill tx pending",
			"request_id": requestId,
			"tx_hash":    request.GetFulfillTxHash(),
		}).Info("tx still pending - ignore")
		return
	}

	// try and get the receipt
	fulfillReceipt, err := d.client.TransactionReceipt(context.Background(), fulfilTxHash)
	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processPossiblyStuck",
			"action":     "get fulfil tx receipt",
			"request_id": request.GetRequestId(),
			"tx_hash":    request.GetFulfillTxHash(),
		}).Error(err.Error())
		return
	}

	if fulfillReceipt.Status == 1 {
		// Tx was successful. Move on and wait for RandomnessRequestFulfilled event
		// to be picked up by the ProcessIncommingEvents function
		// todo - check if the event was missed for some reason and try to get event data
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "processPossiblyStuck",
			"action":     "check fulfill tx status",
			"request_id": requestId,
		}).Info("tx success. wait for RandomnessRequestFulfilled event")
		return
	}

	// Tx has failed - get the reason and try to re-send

	// used later to store failed fulfill tx history
	failedGasUsed := fulfillReceipt.GasUsed
	failedGasPrice := lastFulfillTx.GasPrice().Uint64()
	// todo - try to get actual revert reason
	failReason := "transaction reverted"

	// Add fail info to failed Tx history table
	_ = d.service.Store.Db.InsertNewFailedFulfilment(requestId, request.GetFulfillTxHash(), failedGasUsed, failedGasPrice, failReason)

	// at some point, we just have to stop trying...
	if request.GetFulfillmentAttempts() >= 3 {
		// too many fails
		d.logger.WithFields(logrus.Fields{
			"package":      "chainlisten",
			"function":     "processPossiblyStuck",
			"action":       "check num attempts",
			"request_id":   requestId,
			"num_attempts": request.GetFulfillmentAttempts(),
		}).Info()

		_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, "too many failed attempts")
		return
	}

	// finally, try to send a new fulfillment
	// check if the block hash for the request has been stored in BlockHash contract
	foundHash, _, bhErr := d.service.VORCoordinatorCaller.GetBlockHashFromBlockStore(requestTxReceipt.BlockNumber.Uint64())

	// is the request really old?
	if requestBlockDiff > 250 {
		// check if hash is in block store
		if !foundHash || bhErr != nil {
			// block hash not in store. Flag as failed and move on
			d.logger.WithFields(logrus.Fields{
				"package":    "chainlisten",
				"function":   "processPossiblyStuck",
				"action":     "check request age",
				"request_id": request.GetRequestId(),
			}).Info("request too old and block hash not in block store")
			_ = d.service.Store.Db.UpdateRequestStatus(request.GetRequestId(), database.REQUEST_STATUS_FULFILMENT_FAILED, "request too old and block hash not in block store")
			return
		}
	}

	// record the block hash in the contract to be safe
	if requestBlockDiff > 50 && !foundHash {
		d.recordBlockHash(requestTxReceipt.BlockNumber.Uint64(), requestTxReceipt.BlockHash.Hex())
	}

	d.processFulfillment(requestId, requestTxReceipt, currentBlockNum)

	return
}

func (d *VORCoordinatorListener) preProcessJob(request database.RandomnessRequest, currentBlockNum uint64) {
	requestId := request.GetRequestId()
	d.logger.WithFields(logrus.Fields{
		"package":    "chainlisten",
		"function":   "preProcessJob",
		"action":     "preprocess job",
		"request_id": requestId,
		"status":     request.GetStatusString(),
	}).Info()

	// get request Tx receipt from chain
	requestTxReceipt, err := d.client.TransactionReceipt(context.Background(), common.HexToHash(request.GetRequestTxHash()))
	if err != nil {
		// possibly not in Tx pool yet
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "preProcessJob",
			"action":     "get tx receipt from chain",
			"request_id": requestId,
			"request_tx": request.GetRequestTxHash(),
		}).Error(err.Error())
		return
	}

	requestBlockDiff := currentBlockNum - requestTxReceipt.BlockNumber.Uint64()
	switch request.GetStatus() {
	case database.REQUEST_STATUS_INITIALISED:
		if requestBlockDiff >= config.Conf.WaitConfirmations {
			d.processFulfillment(requestId, requestTxReceipt, currentBlockNum)
		} else {
			// log it
			d.logger.WithFields(logrus.Fields{
				"package":       "chainlisten",
				"function":      "preProcessJob",
				"action":        "check confirmations for initialised job",
				"request_id":    requestId,
				"request_block": requestTxReceipt.BlockNumber.Uint64(),
				"current_block": currentBlockNum,
				"block_diff":    requestBlockDiff,
				"wait_config":   config.Conf.WaitConfirmations,
			}).Info("not enough block confirmations to fulfill request")
		}
		return
	case database.REQUEST_STATUS_TX_FAILED:
		d.processFailed(request, requestTxReceipt, currentBlockNum)
		return
	case database.REQUEST_STATUS_SENT:
		d.processPossiblyStuck(request, requestTxReceipt, currentBlockNum)
		return
	default:
		return
	}
}

func (d *VORCoordinatorListener) CheckJobs() error {

	d.logger.WithFields(logrus.Fields{
		"package":  "chainlisten",
		"function": "CheckJobs",
		"action":   "check job queue",
	}).Info()

	currentBlockNum, err := d.client.BlockNumber(context.Background())

	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "CheckJobs",
			"action":   "get block num",
		}).Error(err.Error())
		return err
	}

	// get requests status = INITIALISED || SENT || FAILED_TX from request_randomness table
	requests, err := d.service.Store.Db.GetJobs()

	if err != nil {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "CheckJobs",
			"action":   "get job requests",
		}).Error(err.Error())
		return err
	}

	for _, request := range requests {
		// process
		d.preProcessJob(request, currentBlockNum)
	}

	return nil
}

func (d *VORCoordinatorListener) ProcessIncommingEvents() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	thisBlockNum, err := d.client.BlockNumber(context.Background())
	if err == nil {
		_ = d.SetLastBlockNumber(thisBlockNum)
	}

	if len(logs) == 0 {
		d.logger.WithFields(logrus.Fields{
			"package":    "chainlisten",
			"function":   "ProcessIncommingEvents",
			"action":     "check events",
			"from_block": d.query.FromBlock.Uint64(),
			"to_block":   thisBlockNum,
		}).Info("no applicable logs")

		return nil
	}

	contractAbi, err := abi.JSON(strings.NewReader(vor_coordinator.VorCoordinatorABI))
	if err != nil {
		return err
	}

	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	logRandomnessRequestFulfilledHash := crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))

	for _, vLog := range logs {
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
			err = contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
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

				// check status and if requests already exists
				reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

				if reqDbRes.ID == 0 {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "ProcessIncommingEvents",
						"action":   "add job to db",
					}).Info("new request")

					_ = d.service.Store.Db.InsertNewRequest(
						common.Bytes2Hex(event.KeyHash[:]),
						event.Sender.Hex(),
						requestId,
						database.REQUEST_STATUS_INITIALISED,
						vLog.TxHash.Hex(),
						gasUsed,
						gasPrice,
						event.Fee.Uint64(),
					)
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
