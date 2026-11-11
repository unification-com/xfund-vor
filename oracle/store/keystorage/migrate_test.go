package keystorage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// Ganache's deterministic accounts 0 and 1.
const (
	key0     = "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
	address0 = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	key1     = "0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1"
	address1 = "0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"

	legacyToken   = "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj"
	newPassphrase = "a new passphrase, long enough"
)

// legacyFixture builds a v0.1.x keystore: each key's CipherPrivate is given in plain hex and encrypted here the
// way v0.1.x did.
func legacyFixture(t *testing.T, token string, keys ...legacyKey) []byte {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.MinCost)
	require.NoError(t, err)
	aesKey := sha256.Sum256([]byte(token))
	block, err := aes.NewCipher(aesKey[:])
	require.NoError(t, err)
	for i := range keys {
		out := make([]byte, aes.BlockSize+len(keys[i].CipherPrivate))
		_, err = rand.Read(out[:aes.BlockSize])
		require.NoError(t, err)
		cipher.NewCFBEncrypter(block, out[:aes.BlockSize]).XORKeyStream(out[aes.BlockSize:], []byte(keys[i].CipherPrivate))
		keys[i].CipherPrivate = base64.URLEncoding.EncodeToString(out)
	}
	data, err := json.Marshal(legacyKeystore{Keys: keys, Hash: string(hash)})
	require.NoError(t, err)
	return data
}

func writeKeystore(t *testing.T, data []byte) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "keystore.json")
	require.NoError(t, os.WriteFile(path, data, 0664))
	return path
}

func open(t *testing.T, path string) *Keystorage {
	t.Helper()
	ks, err := NewKeyStorage(logrus.New(), path)
	require.NoError(t, err)
	return ks
}

func TestMigrateFile(t *testing.T) {
	path := writeKeystore(t, legacyFixture(t, legacyToken,
		legacyKey{Account: "oracle", CipherPrivate: key0, Registered: true, BlockNumber: 13452142},
		legacyKey{Account: "spare", CipherPrivate: strings.TrimPrefix(key1, "0x")},
	))

	_, err := NewKeyStorage(logrus.New(), path)
	assert.ErrorIs(t, err, ErrLegacyKeystore, "the oracle must refuse to start on a v0.1.x keystore")

	migrated, err := MigrateFile(path, legacyToken, newPassphrase)
	require.NoError(t, err)
	assert.Equal(t, []MigratedKey{{Account: "oracle", Address: address0}, {Account: "spare", Address: address1}}, migrated)

	data, err := os.ReadFile(path)
	require.NoError(t, err)
	for _, secret := range []string{key0[2:], key1[2:], legacyToken, newPassphrase} {
		assert.NotContains(t, string(data), secret)
	}
	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
	_, err = os.Stat(path + migratingSuffix)
	assert.True(t, os.IsNotExist(err), "the staging file must be gone")

	// every key, its registration and its block survive, under the new passphrase only
	ks := open(t, path)
	assert.Error(t, ks.Unlock("oracle", legacyToken))
	require.NoError(t, ks.Unlock("oracle", newPassphrase))
	assert.Equal(t, key0, ks.GetSelectedPrivateKey())
	assert.True(t, ks.IsRegistered())
	block, err := ks.GetBlockNumber()
	require.NoError(t, err)
	assert.Equal(t, int64(13452142), block)

	require.NoError(t, ks.Unlock("spare", newPassphrase))
	assert.Equal(t, key1, ks.GetSelectedPrivateKey())
	assert.False(t, ks.IsRegistered())

	_, err = MigrateFile(path, legacyToken, newPassphrase)
	assert.ErrorIs(t, err, ErrAlreadyMigrated)
}

// The v0.1.x keystore the old tests used, written by v0.1.x itself.
func TestMigrateFileReadsARealV01Keystore(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "test_data", "keystore_test_keystore.json"))
	require.NoError(t, err)
	path := writeKeystore(t, data)

	migrated, err := MigrateFile(path, legacyToken, newPassphrase)
	require.NoError(t, err)
	assert.Equal(t, []MigratedKey{{Account: "test", Address: address1}}, migrated)

	ks := open(t, path)
	require.NoError(t, ks.Unlock("test", newPassphrase))
	assert.Equal(t, key1, ks.GetSelectedPrivateKey())
}

// Nothing may touch the old keystore unless the migration succeeds.
func TestMigrateFileLeavesTheKeystoreAloneOnFailure(t *testing.T) {
	legacy := legacyFixture(t, legacyToken, legacyKey{Account: "oracle", CipherPrivate: key0})
	cases := map[string]struct {
		data          []byte
		token, newKey string
		want          string
	}{
		"wrong current key": {legacy, "notthetoken", newPassphrase, "does not unlock"},
		"short new key":     {legacy, legacyToken, "short", "at least 16"},
		"new key reused":    {legacyFixture(t, newPassphrase, legacyKey{Account: "oracle", CipherPrivate: key0}), newPassphrase, newPassphrase, "must differ"},
		"not a key":         {legacyFixture(t, legacyToken, legacyKey{Account: "oracle", CipherPrivate: "0xnothex"}), legacyToken, newPassphrase, "not a valid private key"},
		"no keys":           {[]byte(`{"keys":[],"hash":""}`), legacyToken, newPassphrase, "no keys to migrate"},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			path := writeKeystore(t, c.data)

			_, err := MigrateFile(path, c.token, c.newKey)
			require.Error(t, err)
			assert.Contains(t, err.Error(), c.want)

			after, err := os.ReadFile(path)
			require.NoError(t, err)
			assert.Equal(t, string(c.data), string(after))
			_, err = os.Stat(path + migratingSuffix)
			assert.True(t, os.IsNotExist(err), "no staging file may be left behind")
		})
	}
}
