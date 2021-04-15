package service

import (
	"math/big"
)

func (d *Service) QueryFees(consumer string) (*big.Int, error) {
	return d.VORCoordinatorCaller.QueryFees(consumer)
}
