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
	"oracle/contracts/vor_randomness_request_mock"
	"oracle/service"
	"oracle/tools/vor"
	"strings"
	"sync"
)

type VORRandomnessRequestMockListener struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_randomness_request_mock.VORRandomnessRequestMock
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
	instance, err := vor_randomness_request_mock.NewVORRandomnessRequestMock(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockListener{
		client:          client,
		contractAddress: contractAddress,
		instance:        instance,
		service:         service,
		context:         ctx,
		query: ethereum.FilterQuery{
			FromBlock: big.NewInt(1),
			//ToBlock:   big.NewInt(23),
			Addresses: []common.Address{contractAddress},
		},
	}, err
}

func (d VORRandomnessRequestMockListener) StartPoll() {
	d.wg.Add(1)
	d.wg.Wait()
}

func (d *VORRandomnessRequestMockListener) Request() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(vor_randomness_request_mock.VORRandomnessRequestMockABI))
	if err != nil {
		return err
	}
	logRandomnessRequestSig := []byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)")
	logRandomnessRequestHash := crypto.Keccak256Hash(logRandomnessRequestSig)

	fmt.Println("logRandomnessRequestHash: ", logRandomnessRequestHash)
	fmt.Println("logRandomnessRequestHash hex: ", logRandomnessRequestHash.Hex())
	fmt.Println("logs: ", logs)

	for _, vLog := range logs {
		fmt.Println("----------------------------------------")
		fmt.Println("Log Block Number: ", vLog.BlockNumber)
		fmt.Println("Log Index: ", vLog.Index)
		fmt.Println("vLog.Topics[0].Hex(): ", vLog.Topics[0].Hex())
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			fmt.Println("Log Name: RandomnessRequest")

			//var randomnessRequestEvent contractModel.LogRandomnessRequest
			event := vor_randomness_request_mock.VORRandomnessRequestMockRandomnessRequest{}
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

			tx, err := d.service.Oracle.FulfillRandomness(byteSeed, vLog.BlockHash, int64(vLog.BlockNumber))
			if err != nil {
				return err
			}
			fmt.Println(tx)

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
