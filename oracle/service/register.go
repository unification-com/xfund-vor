package service

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
)

func (d *Oracle) Register(privateKey string, fee int64, providerPaysGas bool) (*types.Transaction, error) {
	var err error
	d.VORCoordinatorCaller, err = chaincall.NewVORCoordinatorCaller(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, big.NewInt(config.Conf.NetworkID), []byte(privateKey))
	if err != nil {
		return nil, err
	}
	return d.VORCoordinatorCaller.RegisterProvingKey(*big.NewInt(fee), providerPaysGas)
}
