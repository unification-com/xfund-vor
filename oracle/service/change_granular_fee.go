package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (d *Service) ChangeGranularFee(consumer common.Address, amount int64) (*types.Transaction, error) {
	return d.VORCoordinatorCaller.ChangeGranularFee(consumer, big.NewInt(amount))
}
