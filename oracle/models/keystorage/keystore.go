package keystorage

import "encoding/json"

// FileVersion is the keystore file format this oracle reads and writes. Version 2 holds each private key as a
// standard go-ethereum v3 (scrypt) keystore. Files from v0.1.x carry no version and encrypt keys with the old
// token scheme; `oracle migrate-keystore` converts them.
const FileVersion = 2

// KeyStorageModel is the keystore file.
type KeyStorageModel struct {
	Version int                   `json:"version"`
	Key     []*KeyStorageKeyModel `json:"keys"`
}

// KeyStorageKeyModel is one account in the keystore.
type KeyStorageKeyModel struct {
	Account string `json:"account"`
	// Address is the key's wallet address, so the operator can tell keys apart without decrypting them.
	Address string `json:"address"`
	// Keystore is the private key as a go-ethereum v3 keystore: the object can be extracted and used with any
	// standard Ethereum tool.
	Keystore    json.RawMessage `json:"keystore"`
	Registered  bool            `json:"registered"`
	BlockNumber int64           `json:"block_number"`
}
