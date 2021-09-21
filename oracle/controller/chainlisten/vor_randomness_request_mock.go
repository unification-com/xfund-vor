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
	"oracle/config"
	"oracle/contracts/vor_coordinator"
	"oracle/contracts/vor_randomness_request_mock"
	"oracle/models/database"
	"oracle/service"
	"oracle/tools/vor"
	"strings"
	"sync"
)

type VORRandomnessRequestMockListener struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_randomness_request_mock.VorRandomnessRequestMock
	query           ethereum.FilterQuery
	wg              *sync.WaitGroup
	service         *service.Service
	context         context.Context
}

func NewVORRandomnessRequestMockListener(contractHexAddress string, ethHostAddress string, service *service.Service, ctx context.Context) (*VORRandomnessRequestMockListener, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(contractHexAddress)
	instance, err := vor_randomness_request_mock.NewVorRandomnessRequestMock(contractAddress, client)
	if err != nil {
		return nil, err
	}

	var lastBlock *big.Int
	lastRequest, err := service.Store.Db.GetLast()
	if blockNumber, _ := service.Store.Keystorage.GetBlockNumber(); blockNumber != 0 {
		lastBlock = big.NewInt(blockNumber)
	} else if lastRequest.GetRequestBlockNumber() != 0 {
		lastBlock = big.NewInt(int64(lastRequest.GetRequestBlockNumber()))
	} else if config.Conf.FirstBlockNumber != 0 {
		lastBlock = big.NewInt(int64(config.Conf.FirstBlockNumber))
	} else {
		lastBlock = big.NewInt(1)
	}

	return &VORRandomnessRequestMockListener{
		client:          client,
		contractAddress: contractAddress,
		instance:        instance,
		service:         service,
		context:         ctx,
		query: ethereum.FilterQuery{
			FromBlock: lastBlock,
			//ToBlock:   big.NewInt(23),
			Addresses: []common.Address{contractAddress},
		},
	}, err
}

func (d VORRandomnessRequestMockListener) StartPoll() {
	d.wg.Add(1)
	d.wg.Wait()
}

func (d *VORRandomnessRequestMockListener) SetLastBlockNumber(blockNumber uint64) (err error) {
	d.query.FromBlock = big.NewInt(int64(blockNumber))
	err = d.service.Store.Keystorage.SetBlockNumber(int64(blockNumber))
	return
}

func (d *VORRandomnessRequestMockListener) Request() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(vor_randomness_request_mock.VorRandomnessRequestMockABI))
	if err != nil {
		return err
	}
	logRandomnessRequestHash := crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	logRandomnessRequestFulfilledHash := crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))

	fmt.Println("logRandomnessRequestHash: ", logRandomnessRequestHash)
	fmt.Println("logRandomnessRequestHash hex: ", logRandomnessRequestHash.Hex())
	fmt.Println("logs: ", logs)
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

		if index == len(logs)-1 {
			err = d.SetLastBlockNumber(vLog.BlockNumber)
		}
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			fmt.Println("Log Name: RandomnessRequest")

			//var randomnessRequestEvent contractModel.LogRandomnessRequest
			event := vor_randomness_request_mock.VorRandomnessRequestMockRandomnessRequest{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data)
			if err != nil {
				return err
			}
			fmt.Println(common.Bytes2Hex(event.KeyHash[:]))
			fmt.Println(*event.Seed)
			fmt.Println(event.Sender.Hex())
			fmt.Println(*event.Fee)
			fmt.Println(common.Bytes2Hex(event.RequestID[:]))
			fmt.Println(vLog.BlockHash)
			fmt.Println(vLog.BlockNumber)
			fmt.Println(event.Raw)

			byteSeed, err := vor.BigToSeed(event.Seed)

			var status int
			fulfilTx, err := d.service.FulfillRandomness(byteSeed, vLog.BlockHash, vLog.BlockNumber)
			fmt.Println(fulfilTx)
			if err != nil {
				fmt.Println(err)
				status = database.REQUEST_STATUS_TX_FAILED
			} else {
				status = database.REQUEST_STATUS_SENT
			}
			//seedHex, err := utils.Uint256ToHex(event.Seed)
			err = d.service.Store.Db.InsertNewRequest(
				common.Bytes2Hex(event.KeyHash[:]),
				event.Sender.Hex(),
				common.Bytes2Hex(event.RequestID[:]),
				status,
				vLog.TxHash.Hex(),
				gasUsed,
				gasPrice,
				event.Fee.Uint64(),
			)
			continue
		case logRandomnessRequestFulfilledHash.Hex():
			fmt.Println("Log Name: RandomnessRequestFulfilled")
			event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
			err := contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data)
			fmt.Println(event)
			if err != nil {
				return err
			}

			err = d.service.Store.Db.UpdateFulfillment(
				common.Bytes2Hex(event.RequestId[:]),
				database.REQUEST_STATUS_SUCCESS,
				event.Output.String(),
				vLog.BlockHash.Hex(),
				vLog.BlockNumber,
				vLog.TxHash.Hex(),
				gasUsed,
				gasPrice,
			)
			continue
		default:
			fmt.Println("vLog: ", vLog)
			continue
		}
	}
	return err
}

func (d VORRandomnessRequestMockListener) RandomnessRequest() {

}
