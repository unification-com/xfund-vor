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
	lastRequest, err := service.Store.RandomnessRequest.Last()
	if blockNumber, _ := service.Store.Keystorage.GetBlockNumber(); blockNumber != 0 {
		lastBlock = big.NewInt(blockNumber + 1)
	} else if lastRequest != nil {
		lastBlock = big.NewInt(int64(lastRequest.GetReqBlockNumber()))
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

	contractAbi, err := abi.JSON(strings.NewReader(string(vor_coordinator.VorCoordinatorABI)))
	if err != nil {
		return err
	}

	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	logRandomnessRequestFulfilledHash := crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))

	fmt.Println("logs: ", logs)

	for index, vLog := range logs {
		fmt.Println("----------------------------------------")
		fmt.Println("Log Block Number: ", vLog.BlockNumber)
		fmt.Println("Log Index: ", vLog.Index)

		txRec, _ := d.client.TransactionReceipt(context.Background(), vLog.TxHash)
		tx, _, _ := d.client.TransactionByHash(context.Background(), vLog.TxHash)

		if index == len(logs)-1 {
			err = d.SetLastBlockNumber(vLog.BlockNumber)
		}
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			fmt.Println("Log Name: RandomnessRequest")

			//var randomnessRequestEvent contractModel.LogRandomnessRequest
			event := vor_coordinator.VorCoordinatorRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				return err
			}
			fmt.Println("event.KeyHash: ", event.KeyHash)
			fmt.Println("d.keyHash: ", d.keyHash)
			if event.KeyHash == d.keyHash {

				fmt.Println("It's request to me =)")

				byteSeed, err := vor.BigToSeed(event.Seed)

				var status string
				tx, err := d.service.FulfillRandomness(byteSeed, vLog.BlockHash, int64(vLog.BlockNumber))
				fmt.Println(tx)
				if err != nil {
					fmt.Println(err)
					status = "failed"
				} else {
					status = "pending"
				}
				seedHex, err := utils.Uint256ToHex(event.Seed)
				err = d.service.Store.RandomnessRequest.InsertNewRequest(
					common.Bytes2Hex(event.KeyHash[:]),
					seedHex, event.Sender.Hex(),
					common.Bytes2Hex(event.RequestID[:]),
					vLog.BlockHash.Hex(),
					vLog.BlockNumber,
					vLog.TxHash.Hex(),
					status,
					txRec.GasUsed,
					tx.GasPrice().Uint64(),
					)
			} else {
				fmt.Println("Looks like it's addressed not to me =(")
			}
			continue
		case logRandomnessRequestFulfilledHash.Hex():
			fmt.Println("Log Name: RandomnessRequestFulfilled")
			event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data)
			fmt.Println(event)
			if err != nil {
				return err
			}

			err = d.service.Store.RandomnessRequest.UpdateFulfillment(
				common.Bytes2Hex(event.RequestId[:]),
				vLog.TxHash.Hex(),
				"success",
				txRec.GasUsed,
				vLog.BlockNumber,
				tx.GasPrice().Uint64(),
				event.Output.String(),
				)
			if err != nil {
				return err
			}

			continue
		default:
			fmt.Println("vLog: ", vLog)
			continue
		}
	}

	return err
}

func (d VORCoordinatorListener) RandomnessRequest() {

}
