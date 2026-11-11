package keystorage

import (
	"crypto/ecdsa"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newKeystore(t *testing.T) (*Keystorage, string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "keystore.json")
	return open(t, path), path
}

func TestKeystorage_GenerateTokenIsRandomAndOnlyForANewKeystore(t *testing.T) {
	ks, _ := newKeystore(t)
	first, err := ks.GenerateToken()
	require.NoError(t, err)
	second, err := ks.GenerateToken()
	require.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile(`^[0-9a-f]{64}$`), first)
	assert.NotEqual(t, first, second)

	_, err = ks.GeneratePrivate("oracle")
	require.NoError(t, err)
	_, err = ks.GenerateToken()
	assert.Error(t, err, "a new token would orphan the keys already encrypted under the current one")
}

func TestKeystorage_KeysSurviveAReopenAndUnlockWithThePassphraseOnly(t *testing.T) {
	ks, path := newKeystore(t)
	token, err := ks.GenerateToken()
	require.NoError(t, err)
	generated, err := ks.GeneratePrivate("oracle")
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))

	reopened := open(t, path)
	assert.True(t, reopened.Exists())
	assert.True(t, reopened.ExistsByUsername("owner"))
	assert.Empty(t, reopened.GetSelectedPrivateKey(), "nothing is decrypted until Unlock")

	assert.Error(t, reopened.Unlock("oracle", "the wrong passphrase"))
	require.NoError(t, reopened.Unlock("oracle", token))
	assert.Equal(t, generated, crypto.PubkeyToAddress(mustParse(t, reopened.GetSelectedPrivateKey()).PublicKey).Hex())
	assert.False(t, reopened.IsAPIKey(token), "the keystore key doesn't open the API")

	require.NoError(t, reopened.Unlock("owner", token))
	assert.Equal(t, key0, reopened.GetSelectedPrivateKey())

	// an account that isn't there falls back to the first key, as before
	require.NoError(t, reopened.Unlock("missing", token))
	assert.Equal(t, generated, crypto.PubkeyToAddress(mustParse(t, reopened.GetSelectedPrivateKey()).PublicKey).Hex())
}

func mustParse(t *testing.T, hexKey string) *ecdsa.PrivateKey {
	t.Helper()
	key, err := ParsePrivateKey(hexKey)
	require.NoError(t, err)
	return key
}

func TestKeystorage_FileHoldsStandardV3KeystoresAndNoSecrets(t *testing.T) {
	ks, path := newKeystore(t)
	token, err := ks.GenerateToken()
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))

	data, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.NotContains(t, string(data), key0[2:])
	assert.NotContains(t, string(data), token)

	var file struct {
		Version int `json:"version"`
		Keys    []struct {
			Address  string `json:"address"`
			Keystore struct {
				Version int `json:"version"`
				Crypto  struct {
					KDF string `json:"kdf"`
				} `json:"crypto"`
			} `json:"keystore"`
		} `json:"keys"`
	}
	require.NoError(t, json.Unmarshal(data, &file))
	assert.Equal(t, 2, file.Version)
	require.Len(t, file.Keys, 1)
	assert.Equal(t, address0, file.Keys[0].Address)
	assert.Equal(t, 3, file.Keys[0].Keystore.Version)
	assert.Equal(t, "scrypt", file.Keys[0].Keystore.Crypto.KDF)
}

func TestKeystorage_AddRefusesBadInput(t *testing.T) {
	ks, _ := newKeystore(t)
	assert.Error(t, ks.AddExisting("owner", key0), "a keystore with no passphrase can't encrypt")

	_, err := ks.GenerateToken()
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))
	assert.Error(t, ks.AddExisting("owner", key1), "account names are unique")

	secret := "0x" + key0[3:] + "zz"
	err = ks.AddExisting("other", secret)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), key0[3:], "errors must not echo the key")
}

// A key whose stored address doesn't match what it decrypts to has been corrupted or tampered with.
func TestKeystorage_UnlockRejectsAMismatchedAddress(t *testing.T) {
	ks, path := newKeystore(t)
	token, err := ks.GenerateToken()
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))

	data, err := os.ReadFile(path)
	require.NoError(t, err)
	tampered := regexp.MustCompile(`"address":"0x[0-9a-fA-F]{40}"`).ReplaceAllString(string(data), `"address":"`+address1+`"`)
	require.NoError(t, os.WriteFile(path, []byte(tampered), 0600))

	assert.Error(t, open(t, path).Unlock("owner", token))
}

func TestKeystorage_StateBelongsToTheUnlockedKey(t *testing.T) {
	ks, path := newKeystore(t)
	token, err := ks.GenerateToken()
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))
	require.NoError(t, ks.AddExisting("oracle", key1))

	_, err = ks.GetBlockNumber()
	assert.Error(t, err, "a locked keystore has no selected key")
	assert.Error(t, ks.SetBlockNumber(1))
	assert.Error(t, ks.SetRegistered())
	assert.False(t, ks.IsRegistered())

	require.NoError(t, ks.Unlock("oracle", token))
	require.NoError(t, ks.SetBlockNumber(13452142))
	require.NoError(t, ks.SetRegistered())

	reopened := open(t, path)
	require.NoError(t, reopened.Unlock("owner", token))
	block, err := reopened.GetBlockNumber()
	require.NoError(t, err)
	assert.Zero(t, block)
	assert.False(t, reopened.IsRegistered())

	require.NoError(t, reopened.Unlock("oracle", token))
	block, err = reopened.GetBlockNumber()
	require.NoError(t, err)
	assert.Equal(t, int64(13452142), block)
	assert.True(t, reopened.IsRegistered())
}

func TestKeystorage_RejectsAnUnknownFormat(t *testing.T) {
	path := writeKeystore(t, []byte(`{"version":3,"keys":[]}`))
	_, err := NewKeyStorage(nil, path)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "newer")

	path = writeKeystore(t, []byte(`{"version":2,"keys":[{"account":"oracle","address":"0x01"}]}`))
	_, err = NewKeyStorage(nil, path)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "incomplete")
}

// The API key is separate from the keystore key, and only its hash is stored.
func TestKeystorage_APIKey(t *testing.T) {
	ks, path := newKeystore(t)
	token, err := ks.GenerateToken()
	require.NoError(t, err)
	require.NoError(t, ks.AddExisting("owner", key0))
	assert.False(t, ks.HasAPIKey())
	assert.False(t, ks.IsAPIKey(""), "no API key set: nothing is accepted")
	assert.Error(t, ks.SetAPIKey("short"))

	apiKey, err := RandomKey()
	require.NoError(t, err)
	require.NoError(t, ks.SetAPIKey(apiKey))

	reopened := open(t, path)
	assert.True(t, reopened.HasAPIKey())
	assert.True(t, reopened.IsAPIKey(apiKey))
	assert.False(t, reopened.IsAPIKey(apiKey[:63]))
	assert.False(t, reopened.IsAPIKey(token))

	data, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.NotContains(t, string(data), apiKey)
}
