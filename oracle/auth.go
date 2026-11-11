package main

import (
	"fmt"
	"os"
	"strings"

	"oracle/store/keystorage"
)

// unlock decrypts account's key with passphrase (from -k), or with a key the operator types when that is
// missing or wrong.
func unlock(keystore *keystorage.Keystorage, account, passphrase string) error {
	var err error
	if passphrase != "" {
		if err = keystore.Unlock(account, passphrase); err == nil {
			return nil
		}
	}
	if !stdinIsTerminal() {
		if err == nil {
			return fmt.Errorf("the keystore key (-k) is missing: %w", errNoTerminal)
		}
		return fmt.Errorf("the keystore key (-k) does not unlock the keystore (%v): %w", err, errNoTerminal)
	}
	fmt.Println("")
	fmt.Println("Let's verify it's you =)")
	fmt.Println("")
	fmt.Println("Please enter the cli/HTTP key, which was provided to you by Oracle")
	_, err = promptSecretUntil("Key: ", func(key string) error {
		return keystore.Unlock(account, key)
	})
	if err == nil {
		fmt.Println("Okay, let's continue...")
	}
	return err
}

// getPasswordFromFileOrFlag reads the key from the file flagValue names, or takes flagValue itself as the key
// when no such file can be opened.
func getPasswordFromFileOrFlag(flagValue string) string {
	if data, err := os.ReadFile(flagValue); err == nil {
		return strings.TrimSpace(string(data))
	}
	return strings.TrimSpace(flagValue)
}
