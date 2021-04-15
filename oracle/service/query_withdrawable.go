package service

import (
	"math/big"
)

func (d *Service) QueryWithdrawableTokens() (*big.Int, error) {
	return d.VORCoordinatorCaller.QueryWithdrawableTokens()
}
