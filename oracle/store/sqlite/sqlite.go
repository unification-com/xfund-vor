package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("oracle.db"), &gorm.Config{})
	return &DB{
		db,
	}, err
}
