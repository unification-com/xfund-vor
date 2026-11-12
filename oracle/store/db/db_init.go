package db

import (
	"errors"
	"fmt"
	stdlog "log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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
	db, err := gorm.Open(sqlite.Open(config.Conf.Database.Storage), gormConfig())
	return &DB{
		db,
	}, err
}

func NewPostgresDb() (*DB, error) {
	cfg := config.Conf
	if cfg.Database.Host == "" || cfg.Database.Port == 0 {
		return nil, nil
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Database, cfg.Database.Password)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig())

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// ChainDB is the database as one chain sees it: queries only reach that chain's records, and inserts are
// made for it.
type ChainDB struct {
	*gorm.DB
	ChainID int64
}

// ForChain scopes the database to the chain with network id chainID.
func (d *DB) ForChain(chainID int64) *ChainDB {
	// a new session, so the condition is applied afresh to every query rather than accumulating
	return &ChainDB{DB: d.DB.Where("chain_id = ?", chainID).Session(&gorm.Session{}), ChainID: chainID}
}

// v01RequestIdIndex made request ids unique across the table; they are unique per chain.
const v01RequestIdIndex = "idx_randomness_requests_request_id"

// chainTables are the tables whose records belong to a chain.
var chainTables = []interface{}{&database.RandomnessRequest{}, &database.FailedFulfilment{}, &database.BlocksStored{}}

// Migrate brings the schema up to date for an oracle serving chainIDs, assigning records from before chain
// ids to their chain.
func (d *DB) Migrate(chainIDs []int64) error {
	migrator := d.Migrator()
	if migrator.HasTable(&database.RandomnessRequest{}) && migrator.HasIndex(&database.RandomnessRequest{}, v01RequestIdIndex) {
		if err := migrator.DropIndex(&database.RandomnessRequest{}, v01RequestIdIndex); err != nil {
			return fmt.Errorf("dropping %s: %w", v01RequestIdIndex, err)
		}
	}
	if err := d.AutoMigrate(append(chainTables, &database.ChainCursor{})...); err != nil {
		return err
	}
	return d.backfillChainID(chainIDs)
}

// backfillChainID assigns records from before chain ids (chain_id 0) to their chain: unambiguous only when
// the oracle serves one chain, as every v0.1 oracle did.
func (d *DB) backfillChainID(chainIDs []int64) error {
	var unassigned int64
	for _, table := range chainTables {
		var n int64
		if err := d.Unscoped().Model(table).Where("chain_id = ?", 0).Count(&n).Error; err != nil {
			return err
		}
		unassigned += n
	}
	if unassigned == 0 {
		return nil
	}
	if len(chainIDs) != 1 {
		return fmt.Errorf("%d records predate chain ids, and the config serves %d chains: start once with only "+
			"the chain this database was used for, to assign them", unassigned, len(chainIDs))
	}
	return d.Transaction(func(tx *gorm.DB) error {
		for _, table := range chainTables {
			if err := tx.Unscoped().Model(table).Where("chain_id = ?", 0).Update("chain_id", chainIDs[0]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// IsNotFound reports whether err means the record looked up doesn't exist, as opposed to the lookup failing.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// gormConfig logs slow queries and errors, but not a lookup finding nothing: the oracle expects that answer
// (a new request, a chain without a saved cursor) and handles it.
func gormConfig() *gorm.Config {
	return &gorm.Config{Logger: logger.New(stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	})}
}
