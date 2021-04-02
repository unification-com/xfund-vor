package postgresql

type RandomnessRequestModel struct {
	ID           int64  `gorm:"primary_key" json:"id"`
	KeyHashHex   string `gorm:"key_hash_hex" json:"key_hash_hex"`
	SeedHex      string `gorm:"seed_hex" json:"seed_hex"`
	SenderHex    string `gorm:"sender_hex" json:"sender_hex"`
	RequestIDHex string `gorm:"request_id_hex" json:"request_id_hex"`
	BlockHashHex string `gorm:"block_hash_hex" json:"block_hash_hex"`
	BlockNumber  uint64 `gorm:"block_number" json:"block_number"`
	Status       string `json:"status"`
}
