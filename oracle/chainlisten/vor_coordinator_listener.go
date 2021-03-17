package chainlisten

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"oracle/contracts/vor_coordinator"
)

type VORCoordinatorListener struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_coordinator.Token
}

func NewVORCoordinatorListener(contractHexAddress string, ethHostAddress string) (*VORCoordinatorListener, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(contractHexAddress)
	instance, err := vor_coordinator.NewToken(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorListener{
		client:          client,
		contractAddress: contractAddress,
		instance:        instance,
	}, err
}

func (d VORCoordinatorListener) RandomnessRequest() {
	//query := ethereum.FilterQuery{Addresses: d.contractAddress}
}
