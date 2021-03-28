package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
)

func (d *Service) Register(account string, privateKey string, fee int64, providerPaysGas bool) (tx *types.Transaction, err error) {
	if d.Store.Keystorage.ExistsByUsername(account) {
		return nil, fmt.Errorf("This account name is already used")
	}

	err = d.Store.Keystorage.AddExisting(account, privateKey)

	if err != nil {
		return
	}
	d.VORCoordinatorCaller, err = chaincall.NewVORCoordinatorCaller(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, big.NewInt(config.Conf.NetworkID), []byte(d.Store.Keystorage.GetSelectedPrivateKey()))
	if err != nil {
		return
	}

	return d.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(fee), providerPaysGas)
}
