package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

var Conf = &Config{
	FirstBlockNumber: 1,
	GasLimit:         500000,
	MaxGasPrice:      150,
	Serve: &Serve{
		Host: "0.0.0.0",
		Port: 8445,
	},
	CheckDuration: 15,
	Keystorage: &Keystorage{
		File: "./keystore.json",
	},
	Database: &Database{
		Dialect: "sqlite",
		Storage: "./oracle.db",
	},
}

type Config struct {
	VORCoordinatorContractAddress string `json:"contract_address"`
	BlockHashStoreContractAddress string `json:"blockhash_store_address"`
	ContractCallerAddress         string `json:"contract_caller_address"`
	MockContractAddress           string `json:"mock_contract_address"`
	EthHTTPHost                   string `json:"eth_http_host"`
	EthWSHost                     string `json:"eth_ws_host"`
	NetworkID                     int64  `json:"network_id"`
	FirstBlockNumber              uint64 `json:"first_block"`
	CheckDuration                 int32  `json:"check_duration"`
	Serve                         *Serve `json:"serve"`
	LogLevel                      string `json:"log_level"`
	GasLimit                      int64  `json:"gas_limit"`
	MaxGasPrice                   int64       `json:"max_gas_price"`
	Keystorage                    *Keystorage `json:"keystorage"`
	Database                      *Database   `json:"database"`
}

func NewConfig(filePath string) (*Config, error) {
	filename, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	jsonErr := json.Unmarshal(data, &config)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &config, err
}
