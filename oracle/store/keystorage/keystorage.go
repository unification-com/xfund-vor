package keystorage

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"

	"oracle/models/keystorage"
	"oracle/utils"
)

// keystoreFileMode: the keystore holds the oracle's private keys, so only its owner may read it.
const keystoreFileMode = 0600

// MinPassphraseLength: should the keystore file leak, scrypt alone doesn't protect a guessable passphrase.
const MinPassphraseLength = 16

// ErrLegacyKeystore: the keystore was written by v0.1.x, whose encryption is too weak to keep using.
var ErrLegacyKeystore = errors.New("its keys use the v0.1.x encryption: convert it with `oracle migrate-keystore`")

// Keystorage is the oracle's keystore: its accounts, each a private key encrypted as a go-ethereum v3 keystore
// under one passphrase, plus each account's registration state and last polled block, and the hash of the
// HTTP API key. Unlock decrypts the one key the oracle signs with; no other key is ever decrypted.
type Keystorage struct {
	log  *logrus.Logger
	path string

	mu         sync.Mutex
	file       *keystorage.KeyStorageModel
	passphrase string
	selected   int    // index of the unlocked key in file.Key, -1 while locked
	privateKey string // the unlocked key, 0x-prefixed hex
}

// NewKeyStorage opens the keystore at filePath, creating an empty one if the file doesn't exist. A keystore
// that exists but can't be read or parsed is an error: carrying on with an empty keystore would send start-up
// into the interactive first run, which loops forever when the oracle runs without a terminal.
func NewKeyStorage(log *logrus.Logger, filePath string) (*Keystorage, error) {
	ks := &Keystorage{log: log, path: filePath, selected: -1}

	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		log.WithFields(logrus.Fields{
			"package":  "keystorage",
			"function": "NewKeyStorage",
			"action":   "creating file",
		}).Info()
		ks.file = emptyKeystore()
		return ks, ks.save()
	}
	if err != nil {
		return nil, fmt.Errorf("reading keystore %s: %w", filePath, err)
	}
	if ks.file, err = parseKeystore(data); err != nil {
		return nil, fmt.Errorf("keystore %s: %w", filePath, err)
	}
	// keystores written by older versions are world-readable
	if err = os.Chmod(filePath, keystoreFileMode); err != nil {
		return nil, err
	}
	return ks, nil
}

func emptyKeystore() *keystorage.KeyStorageModel {
	return &keystorage.KeyStorageModel{Version: keystorage.FileVersion, Key: []*keystorage.KeyStorageKeyModel{}}
}

// parseKeystore reads a keystore file. A v0.1.x keystore with keys is ErrLegacyKeystore; one without keys has
// nothing to protect, so it is read as an empty keystore and written in the current format on the next save.
func parseKeystore(data []byte) (*keystorage.KeyStorageModel, error) {
	var file keystorage.KeyStorageModel
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, fmt.Errorf("not valid JSON (%v): restore it from a backup", err)
	}
	switch {
	case file.Version > keystorage.FileVersion:
		return nil, fmt.Errorf("its format (version %d) is newer than this oracle reads (%d)", file.Version, keystorage.FileVersion)
	case file.Version == keystorage.FileVersion:
		for _, key := range file.Key {
			if !common.IsHexAddress(key.Address) || len(key.Keystore) == 0 {
				return nil, fmt.Errorf("account %q is incomplete: restore it from a backup", key.Account)
			}
		}
		return &file, nil
	case len(file.Key) > 0:
		return nil, ErrLegacyKeystore
	}
	return emptyKeystore(), nil
}

// ParsePrivateKey parses a hex private key, with or without 0x. Its error never includes the input.
func ParsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(utils.RemoveHexPrefix(hexKey))
	if err != nil {
		return nil, errors.New("not a valid private key: expected 64 hex characters")
	}
	return privateKey, nil
}

// Exists reports whether the keystore has any keys.
func (d *Keystorage) Exists() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.file.Key) > 0
}

func (d *Keystorage) ExistsByUsername(account string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.indexOf(account) >= 0
}

func (d *Keystorage) indexOf(account string) int {
	for index, key := range d.file.Key {
		if key.Account == account {
			return index
		}
	}
	return -1
}

// RandomKey returns 32 random bytes as hex: a passphrase or API key nobody can guess.
func RandomKey() (string, error) {
	random := make([]byte, 32)
	if _, err := rand.Read(random); err != nil {
		return "", err
	}
	return hex.EncodeToString(random), nil
}

// GenerateToken starts a new keystore with a random passphrase, which encrypts the keys added next. It refuses
// once the keystore has keys: they are encrypted under its passphrase.
func (d *Keystorage) GenerateToken() (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.file.Key) > 0 {
		return "", errors.New("the keystore already has keys, encrypted under its existing key")
	}
	passphrase, err := RandomKey()
	if err != nil {
		return "", err
	}
	d.passphrase = passphrase
	return passphrase, nil
}

// GeneratePrivate adds a new random key for account and returns its address.
func (d *Keystorage) GeneratePrivate(account string) (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}
	return d.add(account, privateKey)
}

// AddExisting adds a hex private key for account.
func (d *Keystorage) AddExisting(account string, privateKey string) error {
	key, err := ParsePrivateKey(privateKey)
	if err != nil {
		return err
	}
	_, err = d.add(account, key)
	return err
}

// add encrypts privateKey under the keystore passphrase and saves it as account, returning its address.
func (d *Keystorage) add(account string, privateKey *ecdsa.PrivateKey) (string, error) {
	if account == "" {
		account = "autogenerated"
	}
	d.mu.Lock()
	passphrase := d.passphrase
	d.mu.Unlock()
	if passphrase == "" {
		return "", errors.New("the keystore is locked")
	}

	// scrypt takes about a second: don't hold the lock (and stall the poller's block saves) meanwhile
	blob, address, err := encryptV3(privateKey, passphrase)
	if err != nil {
		return "", err
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	if d.indexOf(account) >= 0 {
		return "", fmt.Errorf("account %q is already in the keystore", account)
	}
	d.file.Key = append(d.file.Key, &keystorage.KeyStorageKeyModel{Account: account, Address: address.Hex(), Keystore: blob})
	return address.Hex(), d.saveLocked()
}

// Unlock decrypts account's key with passphrase and makes it the key the oracle signs with; an account that
// isn't in the keystore falls back to the first key.
func (d *Keystorage) Unlock(account, passphrase string) error {
	d.mu.Lock()
	if len(d.file.Key) == 0 {
		d.mu.Unlock()
		return errors.New("the keystore has no keys")
	}
	index := d.indexOf(account)
	if index < 0 {
		index = 0
	}
	key := *d.file.Key[index]
	d.mu.Unlock()

	privateKey, err := decryptV3(key.Keystore, passphrase, common.HexToAddress(key.Address))
	if err != nil {
		return fmt.Errorf("account %q: %w", key.Account, err)
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.passphrase = passphrase
	d.selected = index
	d.privateKey = hexutil.Encode(crypto.FromECDSA(privateKey))
	return nil
}

// GetSelectedPrivateKey returns the unlocked key as 0x-prefixed hex, or "" while the keystore is locked.
func (d *Keystorage) GetSelectedPrivateKey() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.privateKey
}

// SetAPIKey makes key the HTTP API key, keeping only its hash.
func (d *Keystorage) SetAPIKey(key string) error {
	if len(key) < MinPassphraseLength {
		return fmt.Errorf("the API key must be at least %d characters", MinPassphraseLength)
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.file.APIKeySHA256 = sha256Hex(key)
	return d.saveLocked()
}

// HasAPIKey reports whether an HTTP API key is set: without one, the API refuses every request.
func (d *Keystorage) HasAPIKey() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.file.APIKeySHA256 != ""
}

// IsAPIKey reports, in constant time, whether key is the HTTP API key.
func (d *Keystorage) IsAPIKey(key string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.file.APIKeySHA256 != "" && subtle.ConstantTimeCompare([]byte(sha256Hex(key)), []byte(d.file.APIKeySHA256)) == 1
}

func sha256Hex(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// IsRegistered reports whether the unlocked key is registered with the VOR coordinator.
func (d *Keystorage) IsRegistered() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	key, err := d.selectedKey()
	return err == nil && key.Registered
}

// SetRegistered records that the unlocked key is registered with the VOR coordinator.
func (d *Keystorage) SetRegistered() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	key, err := d.selectedKey()
	if err != nil {
		return err
	}
	key.Registered = true
	return d.saveLocked()
}

// GetBlockNumber returns the scan position v0.1 saved with the unlocked key, which seeds a single-chain
// oracle's cursor in the database; the keystore isn't written with it any more.
func (d *Keystorage) GetBlockNumber() (int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	key, err := d.selectedKey()
	if err != nil {
		return 0, err
	}
	return key.BlockNumber, nil
}

func (d *Keystorage) selectedKey() (*keystorage.KeyStorageKeyModel, error) {
	if d.selected < 0 {
		return nil, errors.New("the keystore is locked")
	}
	return d.file.Key[d.selected], nil
}

func (d *Keystorage) save() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.saveLocked()
}

func (d *Keystorage) saveLocked() error {
	data, err := json.Marshal(d.file)
	if err != nil {
		return err
	}
	return utils.WriteFileWithMaxPerms(d.path, data, keystoreFileMode)
}
