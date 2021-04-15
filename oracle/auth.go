package main

import (
	"fmt"
	"io/ioutil"
	"oracle/store/keystorage"
	"os"
	"strings"
)

func inputKey(keystore *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Please enter the cli/HTTP key, which was provided to you by Oracle")
	fmt.Print("Key: ")
	var key string
	fmt.Scanf("%s\n", &key)
	err = keystore.CheckToken(key)
	if err == nil {
		fmt.Println("Okay, let's continue...")
		return
	} else {
		fmt.Println("I'm not sure I can decrypt your keystore with this key.")
		err = inputKey(keystore)
		return
	}
	return
}

func auth(keystore *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Let's verify it's you =)")
	err = inputKey(keystore)
	return
}

func getPasswordFromFileOrFlag(flagValue string) string {
	file, err := os.Open(flagValue)
	password := ""
	if err != nil {
		password = flagValue
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err == nil {
		password = string(data)
	}

	return strings.TrimSpace(password)
}
