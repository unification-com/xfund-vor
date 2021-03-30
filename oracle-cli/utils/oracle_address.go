package utils

import (
	"fmt"
	"os"
)

func OracleAddress() string {
	if os.Getenv("ORACLE_HOST") == "" {
		fmt.Println("Can't find Oracle host info.\nPlease, check if Oracle daemon is running and CLI settings are correct.")
	}
	if Settings.Settings.GetOraclePort() == "" {
		fmt.Println("Can't find Oracle port info.\nPlease, check if Oracle daemon is running and CLI settings are correct.")
	}
	address := fmt.Sprintf("%s:%s", os.Getenv("ORACLE_HOST"), os.Getenv("ORACLE_PORT"))
	return address
}
