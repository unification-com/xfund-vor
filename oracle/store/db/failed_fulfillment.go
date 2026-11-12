package db

import (
	"oracle/models/database"
)

func (d *ChainDB) InsertNewFailedFulfilment(requestId string, txHash string, gasUsed uint64, gasPrice uint64, reason string) (err error) {
	err = d.Create(&database.FailedFulfilment{
		ChainId:    d.ChainID,
		RequestId:  requestId,
		TxHash:     txHash,
		GasUsed:    gasUsed,
		GasPrice:   gasPrice,
		FailReason: reason,
	}).Error
	return
}
