package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"oracle/config"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Conf.Database.Storage), &gorm.Config{})
	return &DB{
		db,
	}, err
}
