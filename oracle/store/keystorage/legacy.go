package keystorage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"oracle/models/keystorage"
)

// The v0.1.x keystore encrypted each private key with AES-CFB under sha256(token), with no key derivation and
// no MAC, and kept a bcrypt hash of the token to check it. The token came from a time-seeded math/rand, so it
// is guessable. This file only reads that format, for `oracle migrate-keystore`.

type legacyKeystore struct {
	Keys []legacyKey `json:"keys"`
	Hash string      `json:"hash"`
}

type legacyKey struct {
	Account       string `json:"account"`
	CipherPrivate string `json:"cipherprivate"`
	Registered    bool   `json:"registered"`
	BlockNumber   int64  `json:"block_number"`
}

// migrateLegacy converts a v0.1.x keystore to the current format, every key re-encrypted under newPassphrase
// and checked to decrypt back to the same key. It does no file I/O.
func migrateLegacy(data []byte, token, newPassphrase string) (*keystorage.KeyStorageModel, error) {
	var legacy legacyKeystore
	if err := json.Unmarshal(data, &legacy); err != nil {
		return nil, fmt.Errorf("not valid JSON: %w", err)
	}
	if len(legacy.Keys) == 0 {
		return nil, errors.New("it has no keys to migrate")
	}
	// The cipher has no MAC: a wrong token would decrypt to a wrong key without any error.
	if bcrypt.CompareHashAndPassword([]byte(legacy.Hash), []byte(token)) != nil {
		return nil, errors.New("the current key does not unlock it")
	}

	migrated := &keystorage.KeyStorageModel{Version: keystorage.FileVersion}
	for _, key := range legacy.Keys {
		hexKey, err := decryptLegacy(key.CipherPrivate, token)
		if err != nil {
			return nil, fmt.Errorf("account %q: %w", key.Account, err)
		}
		privateKey, err := ParsePrivateKey(hexKey)
		if err != nil {
			return nil, fmt.Errorf("account %q: %w", key.Account, err)
		}
		blob, address, err := encryptV3(privateKey, newPassphrase)
		if err != nil {
			return nil, fmt.Errorf("account %q: %w", key.Account, err)
		}
		migrated.Key = append(migrated.Key, &keystorage.KeyStorageKeyModel{
			Account:     key.Account,
			Address:     address.Hex(),
			Keystore:    blob,
			Registered:  key.Registered,
			BlockNumber: key.BlockNumber,
		})
	}
	return migrated, nil
}

// decryptLegacy decrypts one v0.1.x key: base64url(IV || AES-CFB ciphertext), keyed by sha256(token).
func decryptLegacy(cipherText, token string) (string, error) {
	encrypted, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("the encrypted key is not valid base64: %w", err)
	}
	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("the encrypted key is %d bytes, shorter than its %d-byte IV", len(encrypted), aes.BlockSize)
	}
	key := sha256.Sum256([]byte(token))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	iv, data := encrypted[:aes.BlockSize], encrypted[aes.BlockSize:]
	//lint:ignore SA1019 reading the v0.1.x format is the only use; nothing is encrypted this way any more
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(data, data)
	return string(data), nil
}
