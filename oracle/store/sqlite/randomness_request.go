package sqlite

import (
	"oracle/models"
	"oracle/models/sqlite"
)

type RandomnessRequestStore struct {
	db *DB
}

func NewRandomnessRequestStore(db *DB) *RandomnessRequestStore {
	return &RandomnessRequestStore{db: db}
}

func (d *RandomnessRequestStore) Insert(keyHashHex string, seedHex string, senderHex string, requestIDHex string, blockHashHex string, blockNumber uint64, transactionID string, status string) (err error) {
	err = d.db.Create(&sqlite.RandomnessRequestStoreModel{
		KeyHashHex:    keyHashHex,
		SeedHex:       seedHex,
		SenderHex:     senderHex,
		RequestIDHex:  requestIDHex,
		BlockHashHex:  blockHashHex,
		BlockNumber:   blockNumber,
		TransactionID: transactionID,
		Status:        status,
	}).Error
	return
}

func (d RandomnessRequestStore) Migrate() (err error) {
	if !d.db.Migrator().HasTable("randomness_request") {
		err = d.db.Migrator().CreateTable(&sqlite.RandomnessRequestStoreModel{})
	}
	return
}

func (d RandomnessRequestStore) Last() (request models.IRandomnessRequestStoreModel, err error) {
	request = sqlite.RandomnessRequestStoreModel{}
	err = d.db.Last(&request).Error
	return
}
