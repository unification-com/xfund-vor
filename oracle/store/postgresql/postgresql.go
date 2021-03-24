package postgresql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oracle/config"
)

const Timeout = 5

type DB struct {
	*gorm.DB
}

func Dial() (*DB, error) {
	cfg := config.Conf
	if cfg.Database.Host == "" || cfg.Database.Port == 0 {
		return nil, nil
	}

	var pgDB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Database, cfg.Database.Password))
	if err != nil {
		return nil, err
	}

	err = pgDB.Exec("SELECT 1").Error
	if err != nil {
		return nil, err
	}

	return &DB{pgDB}, nil
}
