package utils

import (
	"fmt"
	"os"
)

func OracleAddress() string {
	if os.Getenv("ORACLE_HOST") == "" {
		fmt.Println("Can't find Oracle host in environment variable (ORACLE_HOST).\nPlease, check if Oracle daemon is running (daemon will set env variables automatically).")
	}
	if os.Getenv("ORACLE_PORT") == "" {
		fmt.Println("Can't find Oracle port in environment variable (ORACLE_PORT).\nPlease, check if Oracle daemon is running (daemon will set env variables automatically).")
	}
	address := fmt.Sprintf("%s:%s", os.Getenv("ORACLE_HOST"), os.Getenv("ORACLE_PORT"))
	return address
}
