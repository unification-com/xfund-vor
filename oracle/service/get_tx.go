package service

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func (d *Service) GetTxInfo(txHashStr string) (*types.Transaction, *types.Receipt, error) {

	tx, _, err := d.VORCoordinatorCaller.GetTx(txHashStr)
	if err != nil {
		return nil, nil, err
	}

	txRec, err := d.VORCoordinatorCaller.GetTxReceipt(txHashStr)
	if err != nil {
		return tx, nil, err
	}

	return tx, txRec, nil
}
