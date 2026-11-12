package main

import (
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/config"
	"oracle/store/keystorage"
)

func TestIssueAPIKey(t *testing.T) {
	dir := t.TempDir()
	ks, err := keystorage.NewKeyStorage(log, filepath.Join(dir, "keystore.json"))
	require.NoError(t, err)

	out := filepath.Join(dir, "api_key")
	written := func() string {
		data, err := os.ReadFile(out)
		require.NoError(t, err)
		return strings.TrimSpace(string(data))
	}
	require.NoError(t, issueAPIKey(ks, out))
	first := written()
	assert.True(t, ks.IsAPIKey(first))
	info, err := os.Stat(out)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())

	// re-issuing replaces the key: the old one stops working
	require.NoError(t, issueAPIKey(ks, out))
	assert.NotEqual(t, first, written())
	assert.True(t, ks.IsAPIKey(written()))
	assert.False(t, ks.IsAPIKey(first))
	_, err = os.Stat(out + ".new")
	assert.True(t, os.IsNotExist(err), "no staging file left behind")
}

// Rewriting the keystore under a running oracle would be undone by its next save.
func TestEnsureOracleStopped(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	original := config.Conf.Serve
	t.Cleanup(func() { config.Conf.Serve = original })
	config.Conf.Serve = &config.Serve{Host: "127.0.0.1", Port: int32(listener.Addr().(*net.TCPAddr).Port)}

	assert.Error(t, ensureOracleStopped(), "the port is taken")
	require.NoError(t, listener.Close())
	assert.NoError(t, ensureOracleStopped())
}
