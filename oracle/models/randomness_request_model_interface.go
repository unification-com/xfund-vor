package models

type IRandomnessRequestStoreModel interface {
	GetId() int64
	GetKeyHash() string
	GetSeed() string
	GetSender() string
	GetRequestId() string
	GetReqBlockHash() string
	GetReqBlockNumber() uint64
	GetRequestTxHash() string
	GetStatus() string
}
