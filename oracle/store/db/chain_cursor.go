package db

import (
	"gorm.io/gorm/clause"

	"oracle/models/database"
)

// NextBlock is the next block the chain's event scan reads, 0 when none is saved yet.
func (d *ChainDB) NextBlock() (uint64, error) {
	var cursor database.ChainCursor
	err := d.First(&cursor).Error
	if IsNotFound(err) {
		return 0, nil
	}
	return cursor.NextBlock, err
}

// SetNextBlock saves the next block the chain's event scan reads.
func (d *ChainDB) SetNextBlock(block uint64) error {
	return d.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"next_block", "updated_at"}),
	}).Create(&database.ChainCursor{ChainId: d.ChainID, NextBlock: block}).Error
}
