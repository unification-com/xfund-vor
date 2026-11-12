package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/config"
)

func write(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	require.NoError(t, os.WriteFile(path, []byte(body), 0600))
	return path
}

const coordinator = "0xCeB679DE5fEF4F7c2a0e4cB6A4a0C6e6d5A50ED8"

func TestNewConfig(t *testing.T) {
	conf, err := config.NewConfig(filepath.Join("..", "test_data", "generic_test_config.json"))
	require.NoError(t, err)
	require.Len(t, conf.ChainList(), 1)
}

// A v0.1 config keeps its one chain at the top level; settings it leaves out take their defaults, and a
// setting it gives as zero stays zero.
func TestNewConfigTopLevelChain(t *testing.T) {
	conf, err := config.NewConfig(write(t, `{"network_id":137,"eth_http_host":"https://polygon","contract_address":"`+coordinator+`","max_gas_price":0,"keystorage":{"file":"k.json"}}`))
	require.NoError(t, err)
	chains := conf.ChainList()
	require.Len(t, chains, 1)
	assert.Equal(t, "137", chains[0].Name)
	assert.Equal(t, uint64(10), chains[0].WaitConfirmations, "default")
	assert.Equal(t, uint64(2000), chains[0].EventScanBatchBlocks, "default")
	assert.Zero(t, chains[0].MaxGasPrice, "explicit zero: uncapped")
	assert.Equal(t, int32(8445), conf.Serve.Port, "default")
	assert.Equal(t, "127.0.0.1", conf.Serve.Host, "default: not every interface")
	assert.Equal(t, "./oracle.db", conf.Database.Storage, "default")
}

func TestNewConfigChains(t *testing.T) {
	conf, err := config.NewConfig(write(t, `{"chains":[
		{"name":"polygon","network_id":137,"eth_http_host":"https://polygon","contract_address":"`+coordinator+`","wait_confirmations":5},
		{"network_id":1,"eth_http_host":"https://eth","contract_address":"`+coordinator+`"}
	]}`))
	require.NoError(t, err)
	chains := conf.ChainList()
	require.Len(t, chains, 2)
	assert.Equal(t, uint64(5), chains[0].WaitConfirmations)
	assert.Equal(t, uint64(10), chains[1].WaitConfirmations, "each chain gets the defaults")
	assert.Equal(t, int64(150), chains[1].MaxGasPrice, "default")

	for selector, want := range map[string]int64{"polygon": 137, "137": 137, "1": 1} {
		chain, err := conf.ResolveChain(selector)
		require.NoError(t, err)
		assert.Equal(t, want, chain.NetworkID)
	}
	_, err = conf.ResolveChain("")
	assert.Error(t, err, "two chains: one must be chosen")
	_, err = conf.ResolveChain("qom")
	assert.Error(t, err)
}

func TestNewConfigRejects(t *testing.T) {
	for name, body := range map[string]string{
		"not json":       `{`,
		"no chain":       `{}`,
		"both forms":     `{"network_id":1,"eth_http_host":"h","contract_address":"` + coordinator + `","chains":[{"network_id":137,"eth_http_host":"h","contract_address":"` + coordinator + `"}]}`,
		"duplicate":      `{"chains":[{"network_id":1,"eth_http_host":"h","contract_address":"` + coordinator + `"},{"network_id":1,"eth_http_host":"h2","contract_address":"` + coordinator + `"}]}`,
		"bad contract":   `{"network_id":1,"eth_http_host":"h","contract_address":"0x12"}`,
		"zero batch":     `{"network_id":1,"eth_http_host":"h","contract_address":"` + coordinator + `","event_scan_batch_blocks":0}`,
		"missing config": "",
	} {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "missing.json")
			if body != "" {
				path = write(t, body)
			}
			_, err := config.NewConfig(path)
			assert.Error(t, err)
		})
	}
}
