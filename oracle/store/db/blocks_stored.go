package db

import (
	"oracle/models/database"
)

func (d *ChainDB) InsertNewStoredBlock(blockHash string,
	blockNumber uint64, txHash string) (err error) {
	err = d.Create(&database.BlocksStored{
		ChainId:     d.ChainID,
		BlockHash:   blockHash,
		BlockNumber: blockNumber,
		TxHash:      txHash,
	}).Error
	return
}
