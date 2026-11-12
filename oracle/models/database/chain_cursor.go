package database

import "time"

// ChainCursor is where a chain's event scan resumes: the next block to scan.
type ChainCursor struct {
	ChainId   int64 `gorm:"primaryKey;autoIncrement:false"`
	NextBlock uint64
	UpdatedAt time.Time
}
