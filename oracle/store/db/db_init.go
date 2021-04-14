package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"oracle/config"
	"oracle/models/database"
)

type DB struct {
	*gorm.DB
}

func NewDb() (*DB, error) {
	switch config.Conf.Database.Dialect {
	case "sqlite":
		return NewSqliteDb()
	case "postgres":
		return NewPostgresDb()
	default:
		return nil, errors.New("no db dialect in config")
	}
}

func NewSqliteDb() (*DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Conf.Database.Storage), &gorm.Config{})
	return &DB{
		db,
	}, err
}

func NewPostgresDb() (*DB, error) {
	cfg := config.Conf
	if cfg.Database.Host == "" || cfg.Database.Port == 0 {
		return nil, nil
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Database, cfg.Database.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (d DB) Migrate() (err error) {
	err = d.AutoMigrate(&database.RandomnessRequest{})
	return
}
