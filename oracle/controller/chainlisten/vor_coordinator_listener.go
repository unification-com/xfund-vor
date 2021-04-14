package chainlisten

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
}

func NewVORCoordinatorListener(contractHexAddress string, ethHostAddress string, service *service.Service, ctx context.Context) (*VORCoordinatorListener, error) {
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

	fmt.Println("start polling from block", lastBlock.Uint64())

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
	}, err
}

func (d VORCoordinatorListener) StartPoll() (err error) {
	d.wg.Add(1)
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
		fmt.Println("no applicable logs from block", d.query.FromBlock, "to latest")
		return nil
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(vor_coordinator.VorCoordinatorABI)))
	if err != nil {
		return err
	}

	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	logRandomnessRequestFulfilledHash := crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))

	for index, vLog := range logs {
		fmt.Println("----------------------------------------")
		fmt.Println("Log Block Number: ", vLog.BlockNumber)
		fmt.Println("Log Index: ", vLog.Index)

		gasPrice := uint64(0)
		gasUsed := uint64(0)

		txRec, err := d.client.TransactionReceipt(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need a thread to clean up and gather any data when Tx query fails
			gasUsed = txRec.GasUsed
		} else {
			fmt.Println("TransactionReceipt error: ", err)
		}

		tx, _, err := d.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err == nil {
			// todo - need a thread to clean up and gather any data when Tx query fails
			gasPrice = tx.GasPrice().Uint64()
		} else {
			fmt.Println("TransactionByHash error: ", err)
		}

		if index == len(logs)-1 {
			err = d.SetLastBlockNumber(vLog.BlockNumber)
		}
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			fmt.Println("Log Name: RandomnessRequest")

			event := vor_coordinator.VorCoordinatorRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				fmt.Println(err)
				return err
			}

			if event.KeyHash == d.keyHash {

				fmt.Println("It's a request for me =)")

				byteSeed, err := vor.BigToSeed(event.Seed)

				if err != nil {
					fmt.Println(err)
					return err
				}

				// check status and if requests already exists
				requestId := common.Bytes2Hex(event.RequestID[:])
				reqDbRes, _ := d.service.Store.Db.FindByRequestId(requestId)

				if reqDbRes.ID == 0 {
					fmt.Println("new request")
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
						fmt.Println("error sending fulfil Tx", err)
						_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FAILED, err.Error())
					} else {
						fmt.Println("fulfill tx sent for request", requestId, "in tx", tx.Hash().Hex())
						_ = d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_SENT, "")
					}
				} else {
					fmt.Println("request already in db", "request id", reqDbRes.RequestId, "status", reqDbRes.GetStatusString())
				}
			} else {
				fmt.Println("Looks like it's not addressed to me =(")
			}
			continue
		case logRandomnessRequestFulfilledHash.Hex():
			fmt.Println("Log Name: RandomnessRequestFulfilled")
			event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data)
			requestId := common.Bytes2Hex(event.RequestId[:])
			fmt.Println("confirmed request fulfilment for request", requestId)
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
				fmt.Println(err)
				return err
			}
			continue
		default:
			fmt.Println("event not applicable")
			continue
		}
	}

	return err
}

func (d VORCoordinatorListener) RandomnessRequest() {

}
