package chainlisten

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	"oracle/utils"
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
		logger: logger,
	}, err
}

func (d VORCoordinatorListener) StartPoll() (err error) {
	d.wg.Add(1)

	d.logger.WithFields(logrus.Fields{
		"package":  "chainlisten",
		"function": "StartPoll",
		"action": "begin polling",
		"from_block": d.query.FromBlock.Uint64(),
	}).Info()

	var sleepTime = int32(30)
	if config.Conf.CheckDuration != 0 {
		sleepTime = config.Conf.CheckDuration
	}
	for {
		err = d.Request()
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

func (d *VORCoordinatorListener) Request() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	if len(logs) == 0 {
		d.logger.WithFields(logrus.Fields{
			"package":  "chainlisten",
			"function": "Request",
			"action": "check events",
			"from_block": d.query.FromBlock.Uint64(),
		}).Info("no applicable logs")

		thisBlockNum, err := d.client.BlockNumber(context.Background())
		if err == nil {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "set last block",
				"from_block": thisBlockNum - 1,
			}).Info("no applicable logs")
			_ = d.SetLastBlockNumber(thisBlockNum - 1)
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "get block num",
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
			"package":  "chainlisten",
			"function": "Request",
			"action": "log",
			"block_num": vLog.BlockNumber,
			"log_index": vLog.Index,
		}).Info()

		gasPrice := uint64(0)
		gasUsed := uint64(0)

		txRec, err := d.client.TransactionReceipt(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need a thread to clean up and gather any data when Tx query fails
			gasUsed = txRec.GasUsed
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "get TransactionReceipt",
			}).Error(err.Error())
		}

		tx, _, err := d.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need a thread to clean up and gather any data when Tx query fails
			gasPrice = tx.GasPrice().Uint64()
		} else {
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "get TransactionByHash",
			}).Error(err.Error())
		}

		if index == len(logs)-1 {
			err = d.SetLastBlockNumber(vLog.BlockNumber)
		}
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "check event name",
				"event_name": "RandomnessRequest",
			}).Info("processing event")

			event := vor_coordinator.VorCoordinatorRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "Request",
					"action": "UnpackIntoInterface",
				}).Error(err.Error())
				return err
			}

			if event.KeyHash == d.keyHash {

				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "Request",
					"action": "check event keyhash",
				}).Info("It's a request for me =)")

				byteSeed, err := vor.BigToSeed(event.Seed)

				if err != nil {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "Request",
						"action": "BigToSeed",
					}).Error(err.Error())
					return err
				}

				// check status and if requests already exists
				requestId := common.Bytes2Hex(event.RequestID[:])
				reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

				if reqDbRes.ID == 0 {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "Request",
						"action": "check db for request",
					}).Info("new request")

					seedHex, err := utils.Uint256ToHex(event.Seed)

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

					tx, err := d.service.FulfillRandomness(byteSeed, vLog.BlockHash, int64(vLog.BlockNumber))

					if err != nil {
						d.logger.WithFields(logrus.Fields{
							"package":  "chainlisten",
							"function": "Request",
							"action": "fulfill request",
							"request_id": requestId,
						}).Error(err.Error())
						_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FAILED, err.Error())
					} else {
						d.logger.WithFields(logrus.Fields{
							"package":  "chainlisten",
							"function": "Request",
							"action": "fulfill request",
							"request_id": requestId,
							"tx_hash": tx.Hash().Hex(),
						}).Info("fulfill tx sent")
						_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_SENT, "")
					}
				} else {
					d.logger.WithFields(logrus.Fields{
						"package":  "chainlisten",
						"function": "Request",
						"action": "check db for request",
						"request_id":  reqDbRes.RequestId,
						"status": reqDbRes.GetStatusString(),
					}).Info("request already in db")
				}
			} else {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "Request",
					"action": "check event keyhash",
				}).Info("Looks like it's not addressed to me =(")
			}
			continue
		case logRandomnessRequestFulfilledHash.Hex():
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "check event name",
				"event_name": "RandomnessRequestFulfilled",
			}).Info("processing event")

			event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data)
			requestId := common.Bytes2Hex(event.RequestId[:])
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "check request exists",
				"request_id": requestId,
			}).Info()

			reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

			if reqDbRes.ID != 0 {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "Request",
					"action": "confirm fulfillment",
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
						"function": "Request",
						"action": "UpdateFulfillment",
					}).Error(err.Error())
					return err
				}
			} else {
				d.logger.WithFields(logrus.Fields{
					"package":  "chainlisten",
					"function": "Request",
					"action": "confirm fulfillment",
					"request_id": requestId,
				}).Warning("request id does not exist in db. Probably not mine")
			}
			continue
		default:
			d.logger.WithFields(logrus.Fields{
				"package":  "chainlisten",
				"function": "Request",
				"action": "check event name",
			}).Info("event not applicable")
			continue
		}
	}

	return err
}

func (d VORCoordinatorListener) RandomnessRequest() {

}
