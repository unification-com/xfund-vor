package sqlite

type RandomnessRequestStoreModel struct {
	ID            int64  `gorm:"primary_key" json:"id"`
	KeyHashHex    string `gorm:"key_hash_hex" json:"key_hash_hex"`
	SeedHex       string `gorm:"seed_hex" json:"seed_hex"`
	SenderHex     string `gorm:"sender_hex" json:"sender_hex"`
	RequestIDHex  string `gorm:"request_id_hex" json:"request_id_hex"`
	BlockHashHex  string `gorm:"block_hash_hex" json:"block_hash_hex"`
	BlockNumber   uint64 `gorm:"block_number" json:"block_number"`
	TransactionID string `gorm:"transaction_id" json:"transaction_id"`
	Status        string `json:"status"`
}

func (RandomnessRequestStoreModel) TableName() string {
	return "randomness_request"
}

func (r RandomnessRequestStoreModel) GetID() int64 {
	return r.ID
}

func (r RandomnessRequestStoreModel) GetKeyHashHex() string {
	return r.KeyHashHex
}

func (r RandomnessRequestStoreModel) GetSeedHex() string {
	return r.SeedHex
}

func (r RandomnessRequestStoreModel) GetSenderHex() string {
	return r.SenderHex
}

func (r RandomnessRequestStoreModel) GetRequestIDHex() string {
	return r.RequestIDHex
}

func (r RandomnessRequestStoreModel) GetBlockHashHex() string {
	return r.BlockHashHex
}

func (r RandomnessRequestStoreModel) GetBlockNumber() uint64 {
	return r.BlockNumber
}

func (r RandomnessRequestStoreModel) GetTransactionID() string {
	return r.TransactionID
}

func (r RandomnessRequestStoreModel) GetStatus() string {
	return r.Status
}
