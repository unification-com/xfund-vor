package service

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func (d *Service) SetProviderPaysGas(paysGas bool) (*types.Transaction, error) {
	return d.VORCoordinatorCaller.SetProviderPaysGas(paysGas)
}
