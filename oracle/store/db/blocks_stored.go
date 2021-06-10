package db

import (
	"oracle/models/database"
)

func (d *DB) InsertNewStoredBlock(blockHash string,
	blockNumber uint64, txHash string) (err error) {
	err = d.Create(&database.BlocksStored{
		BlockHash:   blockHash,
		BlockNumber: blockNumber,
		TxHash:      txHash,
	}).Error
	return
}
