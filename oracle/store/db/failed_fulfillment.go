package db

import (
	"oracle/models/database"
)

func (d *DB) InsertNewFailedFulfilment(requestId string, txHash string, reason string) (err error) {
	err = d.Create(&database.FailedFulfilment{
		RequestId:  requestId,
		TxHash:     txHash,
		FailReason: reason,
	}).Error
	return
}
