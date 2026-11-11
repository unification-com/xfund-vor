package keystorage_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/store/keystorage"
)

func TestKeystorage_NewKeyStorageCreatesPrivateFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "keystore.json")

	ks, err := keystorage.NewKeyStorage(Log, path)
	require.NoError(t, err)
	assert.False(t, ks.Exists())

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())

	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	assert.True(t, json.Valid(data), "new keystore must be valid JSON: %s", data)
}

// A keystore that can't be parsed must stop the oracle rather than be treated as empty: an empty keystore
// sends start-up into the interactive first run, which looped forever under systemd (QoM, 2025).
func TestKeystorage_NewKeyStorageRejectsCorruptFile(t *testing.T) {
	cases := map[string]string{
		// what the old non-truncating save left behind: a valid object plus the tail of a longer one
		"stale tail": `{"keys":[{"account":"vor_qom","cipherprivate":"abc","registered":true,"block_number":-1}],"hash":"h"}` + `000"}`,
		"not json":   `garbage`,
	}
	for name, content := range cases {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "keystore.json")
			require.NoError(t, ioutil.WriteFile(path, []byte(content), 0600))

			ks, err := keystorage.NewKeyStorage(Log, path)
			assert.Error(t, err)
			assert.Nil(t, ks)
		})
	}
}

func TestKeystorage_NewKeyStorageTightensPermissions(t *testing.T) {
	path := filepath.Join(t.TempDir(), "keystore.json")
	require.NoError(t, ioutil.WriteFile(path, []byte(`{"keys":[]}`), 0664))
	require.NoError(t, os.Chmod(path, 0664))

	_, err := keystorage.NewKeyStorage(Log, path)
	require.NoError(t, err)

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
}

// Saving a keystore that has shrunk must leave valid JSON: the old save wrote over the file in place without
// truncating, leaving the tail of the longer version behind.
func TestKeystorage_SaveAfterShrinkingLeavesValidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "keystore.json")
	ks, err := keystorage.NewKeyStorage(Log, path)
	require.NoError(t, err)

	_, err = ks.GenerateToken()
	require.NoError(t, err)
	_, err = ks.GeneratePrivate("oracle")
	require.NoError(t, err)
	require.NoError(t, ks.SelectPrivateKey("oracle"))
	require.NoError(t, ks.SetBlockNumber(12345678901))
	require.NoError(t, ks.SetBlockNumber(1))

	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	assert.True(t, json.Valid(data), "keystore must be valid JSON after shrinking: %s", data)

	reopened, err := keystorage.NewKeyStorage(Log, path)
	require.NoError(t, err)
	assert.True(t, reopened.Exists())

	leftovers, err := filepath.Glob(filepath.Join(filepath.Dir(path), ".keystore.json.tmp-*"))
	require.NoError(t, err)
	assert.Empty(t, leftovers, "no temporary files may be left behind")
}

func TestKeystorage_EmptyKeystoreHasNoFirstKey(t *testing.T) {
	ks, err := keystorage.NewKeyStorage(Log, filepath.Join(t.TempDir(), "keystore.json"))
	require.NoError(t, err)

	assert.Nil(t, ks.GetFirst())
	assert.Error(t, ks.SelectPrivateKey("missing"))
}
