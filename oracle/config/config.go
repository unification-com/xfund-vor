package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type Keystorage struct {
	File    string `json:"file"`
	Account string `json:"account"`
}

type Serve struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Storage  string `json:"storage"`
	Dialect  string `json:"dialect"`
}

// ChainConfig is one chain the oracle serves.
type ChainConfig struct {
	// Name selects the chain in admin commands (--chain); it defaults to the network id.
	Name                          string `json:"name"`
	NetworkID                     int64  `json:"network_id"`
	EthHTTPHost                   string `json:"eth_http_host"`
	VORCoordinatorContractAddress string `json:"contract_address"`
	BlockHashStoreContractAddress string `json:"blockhash_store_address"`
	FirstBlockNumber              uint64 `json:"first_block"`
	CheckDuration                 int32  `json:"check_duration"`          // poll interval, seconds
	MaxGasPrice                   int64  `json:"max_gas_price"`           // gwei; 0: uncapped
	LegacyGas                     bool   `json:"legacy_gas"`              // legacy-priced txs even where EIP-1559 is available
	WaitConfirmations             uint64 `json:"wait_confirmations"`      // blocks before a request is fulfilled
	EventScanBatchBlocks          uint64 `json:"event_scan_batch_blocks"` // most blocks per eth_getLogs call
}

// Config is the oracle's config: the chains it serves, and process-wide settings.
type Config struct {
	// A config for one chain may keep the chain's settings at the top level instead of in Chains (the v0.1
	// format).
	ChainConfig
	Chains []ChainConfig `json:"chains"`

	// integration tests only
	ContractCallerAddress string `json:"contract_caller_address"`
	MockContractAddress   string `json:"mock_contract_address"`

	Serve      *Serve      `json:"serve"`
	LogLevel   string      `json:"log_level"`
	Keystorage *Keystorage `json:"keystorage"`
	Database   *Database   `json:"database"`
}

// DefaultChain holds the settings a chain gets when its config leaves them out.
func DefaultChain() ChainConfig {
	return ChainConfig{
		FirstBlockNumber:     1,
		CheckDuration:        15,
		MaxGasPrice:          150,
		WaitConfirmations:    10,
		EventScanBatchBlocks: 2000,
	}
}

func defaults() *Config {
	return &Config{
		ChainConfig: DefaultChain(),
		Serve:       &Serve{Host: "127.0.0.1", Port: 8445},
		Keystorage:  &Keystorage{File: "./keystore.json"},
		Database:    &Database{Dialect: "sqlite", Storage: "./oracle.db"},
	}
}

// Conf is the config the oracle runs with.
var Conf = defaults()

// NewConfig reads and checks the config at filePath. Settings it leaves out keep their defaults, chain by
// chain; one it sets to zero stays zero.
func NewConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	conf := defaults()
	if err = json.Unmarshal(data, conf); err != nil {
		return nil, fmt.Errorf("%s is not valid: %w", filePath, err)
	}
	// each chain is decoded over the defaults, which decoding the list itself would not do
	var chains struct {
		Chains []json.RawMessage `json:"chains"`
	}
	if err = json.Unmarshal(data, &chains); err != nil {
		return nil, err
	}
	conf.Chains = make([]ChainConfig, len(chains.Chains))
	for i, raw := range chains.Chains {
		conf.Chains[i] = DefaultChain()
		if err = json.Unmarshal(raw, &conf.Chains[i]); err != nil {
			return nil, fmt.Errorf("chain %d: %w", i+1, err)
		}
	}
	if err = conf.validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", filePath, err)
	}
	return conf, nil
}

// ChainList is the chains the oracle serves, named.
func (c *Config) ChainList() []ChainConfig {
	chains := c.Chains
	if len(chains) == 0 {
		chains = []ChainConfig{c.ChainConfig}
	}
	named := make([]ChainConfig, len(chains))
	for i, chain := range chains {
		if chain.Name == "" {
			chain.Name = strconv.FormatInt(chain.NetworkID, 10)
		}
		named[i] = chain
	}
	return named
}

// ResolveChain finds a chain by name or network id; with one chain, "" selects it.
func (c *Config) ResolveChain(selector string) (ChainConfig, error) {
	chains := c.ChainList()
	if selector == "" {
		if len(chains) == 1 {
			return chains[0], nil
		}
		return ChainConfig{}, fmt.Errorf("the oracle serves %d chains: choose one with --chain (%s)", len(chains), chainNames(chains))
	}
	for _, chain := range chains {
		if chain.Name == selector || strconv.FormatInt(chain.NetworkID, 10) == selector {
			return chain, nil
		}
	}
	return ChainConfig{}, fmt.Errorf("no chain %q: the oracle serves %s", selector, chainNames(chains))
}

func chainNames(chains []ChainConfig) string {
	names := ""
	for i, chain := range chains {
		if i > 0 {
			names += ", "
		}
		names += fmt.Sprintf("%s (%d)", chain.Name, chain.NetworkID)
	}
	return names
}

func (c *Config) validate() error {
	top := c.ChainConfig
	if len(c.Chains) > 0 && (top.NetworkID != 0 || top.EthHTTPHost != "" || top.VORCoordinatorContractAddress != "") {
		return errors.New("chain settings are both at the top level and in \"chains\": use one or the other")
	}
	var errs []error
	names, ids := map[string]bool{}, map[int64]bool{}
	for _, chain := range c.ChainList() {
		if err := chain.validate(); err != nil {
			errs = append(errs, fmt.Errorf("chain %s: %w", chain.Name, err))
		}
		if names[chain.Name] || ids[chain.NetworkID] {
			errs = append(errs, fmt.Errorf("chain %s (%d) is listed twice", chain.Name, chain.NetworkID))
		}
		names[chain.Name], ids[chain.NetworkID] = true, true
	}
	if c.Serve == nil || c.Serve.Port <= 0 {
		errs = append(errs, errors.New("serve.port must be set"))
	}
	if c.Keystorage == nil || c.Keystorage.File == "" {
		errs = append(errs, errors.New("keystorage.file must be set"))
	}
	return errors.Join(errs...)
}

func (c ChainConfig) validate() error {
	var errs []error
	if c.NetworkID <= 0 {
		errs = append(errs, errors.New("network_id must be set"))
	}
	if c.EthHTTPHost == "" {
		errs = append(errs, errors.New("eth_http_host must be set"))
	}
	if !common.IsHexAddress(c.VORCoordinatorContractAddress) {
		errs = append(errs, errors.New("contract_address must be an address"))
	}
	// optional: only old requests' retries consult the block hash store
	if c.BlockHashStoreContractAddress != "" && !common.IsHexAddress(c.BlockHashStoreContractAddress) {
		errs = append(errs, errors.New("blockhash_store_address must be an address"))
	}
	if c.EventScanBatchBlocks == 0 {
		errs = append(errs, errors.New("event_scan_batch_blocks must be above zero"))
	}
	if c.CheckDuration <= 0 {
		errs = append(errs, errors.New("check_duration must be above zero"))
	}
	return errors.Join(errs...)
}
