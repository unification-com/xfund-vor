package database

import "gorm.io/gorm"

type BlocksStored struct {
	gorm.Model
	ChainId     int64  `gorm:"not null;default:0;index"`
	BlockHash   string `gorm:"index"`
	BlockNumber uint64 `gorm:"index"`
	TxHash      string `gorm:"index"`
}

func (BlocksStored) TableName() string {
	return "blocks_stored"
}

func (f BlocksStored) GetId() uint {
	return f.ID
}

func (f BlocksStored) GetBlockHash() string {
	return f.BlockHash
}

func (f BlocksStored) GetBlockNumber() uint64 {
	return f.BlockNumber
}

func (f BlocksStored) GetTxHash() string {
	return f.TxHash
}
