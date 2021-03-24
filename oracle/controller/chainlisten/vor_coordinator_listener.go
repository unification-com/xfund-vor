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
	"oracle/contracts/vor_coordinator"
	"strings"
	"sync"
)

type VORCoordinatorListener struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_coordinator.VORCoordinator
	query           ethereum.FilterQuery
	wg              *sync.WaitGroup
}

func NewVORCoordinatorListener(contractHexAddress string, ethHostAddress string) (*VORCoordinatorListener, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(contractHexAddress)
	instance, err := vor_coordinator.NewVORCoordinator(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorListener{
		client:          client,
		contractAddress: contractAddress,
		instance:        instance,
		query: ethereum.FilterQuery{
			FromBlock: big.NewInt(1),
			//ToBlock:   big.NewInt(23),
			//Addresses: []common.Address{contractAddress},
		},
	}, err
}

func (d VORCoordinatorListener) StartPoll() {
	d.wg.Add(1)
	d.wg.Wait()
}

func (d *VORCoordinatorListener) Request() error {
	logs, err := d.client.FilterLogs(context.Background(), d.query)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(vor_coordinator.VORCoordinatorABI)))
	if err != nil {
		return err
	}
	logRandomnessRequestSig := []byte("randomnessRequest(bytes32,uint256,uint256)")
	logRandomnessRequestHash := crypto.Keccak256Hash(logRandomnessRequestSig)

	fmt.Println("logRandomnessRequestHash: ", logRandomnessRequestHash)
	fmt.Println("logRandomnessRequestHash hex: ", logRandomnessRequestHash.Hex())
	fmt.Println("logs: ", logs)

	for _, vLog := range logs {
		fmt.Println("Log Block Number: ", vLog.BlockNumber)
		fmt.Println("Log Index: ", vLog.Index)
		fmt.Println("vLog.Topics[0].Hex(): ", vLog.Topics[0].Hex())
		switch vLog.Topics[0].Hex() {
		case logRandomnessRequestHash.Hex():
			fmt.Println("Log Name: randomnessRequest")

			//var randomnessRequestEvent contractModel.LogRandomnessRequest

			data, err := contractAbi.Unpack("randomnessRequest", vLog.Data)
			fmt.Println(data)
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
