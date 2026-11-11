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
	key, err := issueAPIKey(ks, out)
	require.NoError(t, err)
	assert.True(t, ks.IsAPIKey(key))
	written, err := os.ReadFile(out)
	require.NoError(t, err)
	assert.Equal(t, key, strings.TrimSpace(string(written)))
	info, err := os.Stat(out)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())

	// an existing file is never overwritten, and the key in use stays as it was
	_, err = issueAPIKey(ks, out)
	assert.Error(t, err)
	assert.True(t, ks.IsAPIKey(key))
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
