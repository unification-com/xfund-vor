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

type Config struct {
	VORCoordinatorContractAddress string      `json:"contract_address"`
	EthHTTPHost                   string      `json:"eth_http_host"`
	EthWSHost                     string      `json:"eth_ws_host"`
	NetworkID                     int64       `json:"network_id"`
	Serve                         *Serve      `json:"serve"`
	LogFile                       string      `json:"log_file"`
	LogLevel                      string      `json:"log_level"`
	Keystorage                    *Keystorage `json:"keystorage"`
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
