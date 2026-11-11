package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"oracle/store/keystorage"
	"os"
	"strings"
)

// auth asks the operator for the keystore key when the password file is missing or holds the wrong key.
func auth(keystore *keystorage.Keystorage) error {
	if !stdinIsTerminal() {
		return fmt.Errorf("the keystore password (-k) is missing or wrong: %w", errNoTerminal)
	}
	fmt.Println("")
	fmt.Println("Let's verify it's you =)")
	fmt.Println("")
	fmt.Println("Please enter the cli/HTTP key, which was provided to you by Oracle")
	_, err := promptUntil("Key: ", func(key string) error {
		if keystore.CheckToken(key) != nil {
			return errors.New("I'm not sure I can decrypt your keystore with this key")
		}
		return nil
	})
	if err == nil {
		fmt.Println("Okay, let's continue...")
	}
	return err
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
