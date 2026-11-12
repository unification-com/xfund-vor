package keystorage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"oracle/utils"
)

// ErrAlreadyMigrated: MigrateFile found a keystore that is already in the current format.
var ErrAlreadyMigrated = errors.New("the keystore is already in the current format")

// migratingSuffix names the converted keystore while it is written and verified, before it replaces the old one.
const migratingSuffix = ".migrating"

// MigratedKey is an account MigrateFile converted.
type MigratedKey struct {
	Account string
	Address string
}

// MigrateFile converts the v0.1.x keystore at path to the current format, re-encrypting every key under
// newPassphrase. The converted keystore is written beside the old one and every key in it is proved, from the
// bytes on disk, to decrypt and sign as its address before the old file is touched; any failure before then
// leaves the old file as it was. The old file is then overwritten and replaced rather than kept, since a copy
// encrypted under the old token keeps the weakness being removed: back the directory up first.
func MigrateFile(path, token, newPassphrase string) ([]MigratedKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading keystore %s: %w", path, err)
	}
	switch current, err := parseKeystore(data); {
	case errors.Is(err, ErrLegacyKeystore):
	case err != nil:
		return nil, fmt.Errorf("keystore %s: %w", path, err)
	case len(current.Key) > 0:
		return nil, ErrAlreadyMigrated
	default:
		return nil, fmt.Errorf("keystore %s has no keys to migrate", path)
	}
	if err = CheckNewPassphrase(newPassphrase, token); err != nil {
		return nil, err
	}

	migrated, err := migrateLegacy(data, token, newPassphrase)
	if err != nil {
		return nil, fmt.Errorf("keystore %s: %w", path, err)
	}
	converted, err := json.Marshal(migrated)
	if err != nil {
		return nil, err
	}

	staging := path + migratingSuffix
	if err = utils.WriteFileWithMaxPerms(staging, converted, keystoreFileMode); err != nil {
		return nil, fmt.Errorf("writing the converted keystore: %w", err)
	}
	if err = verifyFile(staging, newPassphrase); err != nil {
		_ = os.Remove(staging)
		return nil, fmt.Errorf("the converted keystore failed verification, so %s is unchanged: %w", path, err)
	}

	overwrite(path)
	if err = utils.RenameDurably(staging, path); err != nil {
		return nil, fmt.Errorf("the converted keystore is verified at %s but could not replace %s (move it there "+
			"by hand): %w", staging, path, err)
	}

	keys := make([]MigratedKey, 0, len(migrated.Key))
	for _, key := range migrated.Key {
		keys = append(keys, MigratedKey{Account: key.Account, Address: key.Address})
	}
	return keys, nil
}

// CheckNewPassphrase vets a new keystore passphrase: long enough to resist guessing, and not the v0.1.x token
// it replaces, which came from a guessable random source (and doubled as the API key).
func CheckNewPassphrase(newPassphrase, token string) error {
	if len(newPassphrase) < MinPassphraseLength {
		return fmt.Errorf("the new key must be at least %d characters", MinPassphraseLength)
	}
	if newPassphrase == token {
		return errors.New("the new key must differ from the current one")
	}
	return nil
}

// verifyFile re-reads the keystore at path and proves every key in it decrypts under passphrase and signs as
// its address.
func verifyFile(path, passphrase string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	file, err := parseKeystore(data)
	if err != nil {
		return err
	}
	for _, key := range file.Key {
		if err = verifyV3(key.Keystore, passphrase, common.HexToAddress(key.Address)); err != nil {
			return fmt.Errorf("account %q: %w", key.Account, err)
		}
	}
	return nil
}

// overwrite zeroes a file's contents so they don't linger on disk once it is replaced. Best effort: SSDs and
// copy-on-write filesystems may keep other copies.
func overwrite(path string) {
	info, err := os.Stat(path)
	if err != nil || !info.Mode().IsRegular() {
		return
	}
	f, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = f.WriteAt(make([]byte, info.Size()), 0); err == nil {
		_ = f.Sync()
	}
}
