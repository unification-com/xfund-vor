package database

import "gorm.io/gorm"

const (
	REQUEST_STATUS_UNKNOWN = iota
	REQUEST_STATUS_INITIALISED
	REQUEST_STATUS_SENT
	REQUEST_STATUS_FAILED
	REQUEST_STATUS_SUCCESS
)

type RandomnessRequest struct {
	gorm.Model
	KeyHash            string
	Seed               string
	Sender             string `gorm:"index"`
	RequestId          string `gorm:"uniqueIndex"`
	RequestBlockHash   string `gorm:"index"`
	RequestBlockNumber uint64 `gorm:"index"`
	RequestTxHash      string `gorm:"index"`
	RequestGasUsed     uint64
	RequestGasPrice    uint64
	Fee                uint64
	Randomness         string
	FulfillBlockHash   string `gorm:"index"`
	FulfillBlockNumber uint64 `gorm:"index"`
	FulfillTxHash      string `gorm:"index"`
	FulfillGasUsed     uint64
	FulfillGasPrice    uint64
	Status             int    `gorm:"index"`
	StatusReason       string
}

func (RandomnessRequest) TableName() string {
	return "randomness_requests"
}

func (r RandomnessRequest) GetId() uint {
	return r.ID
}

func (r RandomnessRequest) GetKeyHash() string {
	return r.KeyHash
}

func (r RandomnessRequest) GetSeed() string {
	return r.Seed
}

func (r RandomnessRequest) GetSender() string {
	return r.Sender
}

func (r RandomnessRequest) GetRequestId() string {
	return r.RequestId
}

func (r RandomnessRequest) GetRequestBlockHash() string {
	return r.RequestBlockHash
}

func (r RandomnessRequest) GetRequestBlockNumber() uint64 {
	return r.RequestBlockNumber
}

func (r RandomnessRequest) GetRequestTxHash() string {
	return r.RequestTxHash
}

func (r RandomnessRequest) GetRequestGasUsed() uint64 {
	return r.RequestGasUsed
}

func (r RandomnessRequest) GetRequestGasPrice() uint64 {
	return r.RequestGasPrice
}

func (r RandomnessRequest) GetFee() uint64 {
	return r.Fee
}

func (r RandomnessRequest) GetRandomness() string {
	return r.Randomness
}

func (r RandomnessRequest) GetFulfillBlockHash() string {
	return r.FulfillBlockHash
}

func (r RandomnessRequest) GetFulfillBlockNumber() uint64 {
	return r.FulfillBlockNumber
}

func (r RandomnessRequest) GetFulfillTxHash() string {
	return r.FulfillTxHash
}

func (r RandomnessRequest) GetFulfillGasUsed() uint64 {
	return r.FulfillGasUsed
}

func (r RandomnessRequest) GetFulfillGasPrice() uint64 {
	return r.FulfillGasPrice
}

func (r RandomnessRequest) GetStatus() int {
	return r.Status
}

func (r RandomnessRequest) GetStatusString() string {
	switch r.Status {
	case REQUEST_STATUS_UNKNOWN:
	default:
		return "UNKNOWN"
	case REQUEST_STATUS_INITIALISED:
		return "INITIALISED"
	case REQUEST_STATUS_SENT:
		return "SENT"
	case REQUEST_STATUS_FAILED:
		return "FAILED"
	case REQUEST_STATUS_SUCCESS:
		return "SUCCESS"
	}

	return "UNKNOWN"
}

func (r RandomnessRequest) GetStatusReason() string {
	return r.StatusReason
}
