package store

import "oracle/models"

type IRandomnessRequestStore interface {
	Insert(keyHashHex string, seedHex string, senderHex string, requestIDHex string, blockHashHex string, blockNumber uint64, transactionID string, status string) error
	Migrate() (err error)
	Last() (request models.IRandomnessRequestStoreModel, err error)
}
