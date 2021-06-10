package database

import "gorm.io/gorm"

type FailedFulfilment struct {
	gorm.Model
	RequestId   string `gorm:"index"`
	TxHash      string `gorm:"index"`
	FailReason  string
}

func (FailedFulfilment) TableName() string {
	return "failed_fulfilments"
}

func (f FailedFulfilment) GetId() uint {
	return f.ID
}

func (f FailedFulfilment) GetRequestId() string {
	return f.RequestId
}

func (f FailedFulfilment) GetTxHash() string {
	return f.TxHash
}

func (f FailedFulfilment) GetFailReason() string {
	return f.FailReason
}
