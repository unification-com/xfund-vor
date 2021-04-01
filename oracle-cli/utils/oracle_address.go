package utils

import (
	"fmt"
)

func OracleAddress() string {
	if Settings.Settings.GetOracleHost() == "" {
		fmt.Println("Can't find Oracle host info.\nPlease, check if Oracle daemon is running and CLI settings are correct.")
	}

	if Settings.Settings.GetOraclePort() == "" {
		fmt.Println("Can't find Oracle port info.\nPlease, check if Oracle daemon is running and CLI settings are correct.")
	}
	address := fmt.Sprintf("http://%s:%s", Settings.Settings.GetOracleHost(), Settings.Settings.GetOraclePort())
	return address
}
