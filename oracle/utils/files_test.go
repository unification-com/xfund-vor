package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteFileWithMaxPerms(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "keystore.json")

	require.NoError(t, WriteFileWithMaxPerms(path, []byte(`{"a longer first version"}`), 0600))
	require.NoError(t, WriteFileWithMaxPerms(path, []byte(`{"short"}`), 0600))

	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, `{"short"}`, string(data), "a shorter rewrite must not leave the old tail behind")

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())

	entries, err := ioutil.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, entries, 1, "no temporary files may be left behind")
}
