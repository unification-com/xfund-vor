package models

type IRandomnessRequestStoreModel interface {
	GetID() int64
	GetKeyHashHex() string
	GetSeedHex() string
	GetSenderHex() string
	GetRequestIDHex() string
	GetBlockHashHex() string
	GetBlockNumber() uint64
	GetTransactionID() string
	GetStatus() string
}
