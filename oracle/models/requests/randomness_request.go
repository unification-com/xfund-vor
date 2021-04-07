package requests

type RandomnessRequestStoreModel struct {
	Id              int64  `gorm:"primary_key" json:"id"`
	KeyHash         string `gorm:"key_hash" json:"key_hash"`
	Seed            string `gorm:"seed" json:"seed"`
	Sender          string `gorm:"sender" json:"sender"`
	RequestId       string `gorm:"request_id" json:"request_id"`
	ReqBlockHash    string `gorm:"req_block_hash" json:"req_block_hash"`
	ReqBlockNumber  uint64 `gorm:"req_block_number" json:"req_block_number"`
	RequestTxHash   string `gorm:"req_tx_hash" json:"req_tx_hash"`
	Randomness      string `gorm:"randomness" json:"randomness"`
	FulfillTxHash   string `gorm:"fulfill_tx_hash" json:"fulfill_tx_hash"`
	FulfillBlockNum uint64 `gorm:"fulfill_block_num" json:"fulfill_block_num"`
	FulfillGasUsed  uint64 `gorm:"fulfill_gas_used" json:"fulfill_gas_used"`
	Status          string `gorm:"status" json:"status"`
	StatusReason    string `gorm:"status_reason" json:"status_reason"`
}

func (RandomnessRequestStoreModel) TableName() string {
	return "randomness_request"
}

func (r RandomnessRequestStoreModel) GetId() int64 {
	return r.Id
}

func (r RandomnessRequestStoreModel) GetKeyHash() string {
	return r.KeyHash
}

func (r RandomnessRequestStoreModel) GetSeed() string {
	return r.Seed
}

func (r RandomnessRequestStoreModel) GetSender() string {
	return r.Sender
}

func (r RandomnessRequestStoreModel) GetRequestId() string {
	return r.RequestId
}

func (r RandomnessRequestStoreModel) GetReqBlockHash() string {
	return r.ReqBlockHash
}

func (r RandomnessRequestStoreModel) GetReqBlockNumber() uint64 {
	return r.ReqBlockNumber
}

func (r RandomnessRequestStoreModel) GetRequestTxHash() string {
	return r.RequestTxHash
}

func (r RandomnessRequestStoreModel) GetStatus() string {
	return r.Status
}
