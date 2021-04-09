package store

import "oracle/models"

type IRandomnessRequestStore interface {
	InsertNewRequest(keyHashHex string, seedHex string, senderHex string, requestIDHex string, blockHashHex string, blockNumber uint64, transactionID string, status string, gasUsed uint64, gasPrice uint64) error
	Migrate() (err error)
	Last() (request models.IRandomnessRequestStoreModel, err error)
	UpdateFulfillment(requestId string, fulfillTxHash string, status string, gasUsed uint64, blockNum uint64, gasPrice uint64, randomness string) error
}
