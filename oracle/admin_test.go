package main

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type apiCall struct {
	method, path, query, auth string
	body                      map[string]interface{}
}

// The admin commands find the oracle from its config, authenticate with the api_key file beside it, and
// send what the oracle's handlers expect.
func TestAdminCommands(t *testing.T) {
	var calls []apiCall
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := apiCall{method: r.Method, path: r.URL.Path, query: r.URL.RawQuery, auth: r.Header.Get("Authorization")}
		if b, _ := io.ReadAll(r.Body); len(b) > 0 {
			_ = json.Unmarshal(b, &call.body)
		}
		calls = append(calls, call)
		if r.URL.Path == "/changefee" {
			http.Error(w, "fee must be above zero", http.StatusBadRequest)
			return
		}
		_, _ = w.Write([]byte("1500000000"))
	}))
	defer server.Close()
	host, port, err := net.SplitHostPort(server.Listener.Addr().String())
	require.NoError(t, err)

	dir := t.TempDir()
	configPath := filepath.Join(dir, "config.json")
	portNum, _ := strconv.Atoi(port)
	require.NoError(t, os.WriteFile(configPath, []byte(`{"serve":{"host":"`+host+`","port":`+strconv.Itoa(portNum)+`}}`), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "api_key"), []byte("the-api-key\n"), 0600))

	run := func(args ...string) error {
		root := rootCommand()
		root.SetArgs(append(args, "-c", configPath))
		return root.Execute()
	}

	require.NoError(t, run("withdraw", "250", "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"))
	require.NoError(t, run("queryfees", "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"))
	require.NoError(t, run("queryrequests", "--limit", "20", "--status", "3"))
	err = run("changefee", "5")
	require.Error(t, err, "the oracle's refusal is an error")
	assert.Contains(t, err.Error(), "fee must be above zero")
	assert.Error(t, run("withdraw", "lots", "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"), "amounts are checked before calling")

	require.Len(t, calls, 4)
	for _, call := range calls {
		assert.Equal(t, "Bearer the-api-key", call.auth)
	}
	assert.Equal(t, apiCall{method: "POST", path: "/withdraw", auth: "Bearer the-api-key",
		body: map[string]interface{}{"address": "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1", "amount": float64(250)}}, calls[0])
	assert.Equal(t, "/queryfees", calls[1].path)
	assert.Equal(t, "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1", calls[1].body["consumer"])
	assert.Equal(t, "limit=20&order=desc&page=1&status=3", calls[2].query)
}

func TestAdminCommandsNeedTheAPIKey(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "config.json")
	require.NoError(t, os.WriteFile(configPath, []byte(`{"serve":{"host":"0.0.0.0","port":1}}`), 0600))
	root := rootCommand()
	root.SetArgs([]string{"status", "-c", configPath})
	err := root.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "issue one with `oracle api-key`")
}
