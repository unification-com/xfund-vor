package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"oracle/config"
	"oracle/store/keystorage"
)

// migrateKeystore converts a v0.1.x keystore to the current format: every key re-encrypted as a go-ethereum v3
// (scrypt) keystore under a new key, which replaces the old one for -k. The old key also opened the HTTP API;
// the API now has its own key (`oracle api-key`).
//
//	oracle migrate-keystore -c config.json -k password [--new-key password.new]
func migrateKeystore() error {
	if err := ensureOracleStopped(); err != nil {
		return err
	}
	path := config.Conf.Keystorage.File
	current := ""
	if options.PasswordFile != "" {
		current = getPasswordFromFileOrFlag(options.PasswordFile)
	}
	if current == "" {
		if !stdinIsTerminal() {
			return fmt.Errorf("the current key (-k) is missing: %w", errNoTerminal)
		}
		var err error
		if current, err = promptSecret("Current keystore key: "); err != nil {
			return err
		}
	}
	newKey, err := newKeystoreKey(current)
	if err != nil {
		return err
	}

	migrated, err := keystorage.MigrateFile(path, current, newKey)
	if errors.Is(err, keystorage.ErrAlreadyMigrated) {
		fmt.Println("The keystore", path, "is already in the current format: nothing to do.")
		return nil
	}
	if err != nil {
		return err
	}

	fmt.Println("Migrated", path)
	for _, key := range migrated {
		fmt.Printf("  %s  %s\n", key.Address, key.Account)
	}
	fmt.Println("The old key no longer unlocks anything. Next:")
	fmt.Println("  1. put the new key where -k reads it")
	fmt.Println("  2. issue an API key for oracle-cli: oracle api-key -c <config>")
	fmt.Println("  3. delete any copy of the old keystore and its key, including backups: together they expose the keys")
	return nil
}

// newKeystoreKey reads the new key from the --new-key file, or asks for it twice on a terminal. Unlike -k,
// --new-key must name a file: a mistyped path must not silently become the key.
func newKeystoreKey(current string) (string, error) {
	check := func(key string) error { return keystorage.CheckNewPassphrase(key, current) }
	if options.NewKeyFile != "" {
		data, err := os.ReadFile(options.NewKeyFile)
		if err != nil {
			return "", fmt.Errorf("reading the new key: %w", err)
		}
		key := strings.TrimSpace(string(data))
		return key, check(key)
	}
	if !stdinIsTerminal() {
		return "", fmt.Errorf("no file holds the new key (--new-key): %w", errNoTerminal)
	}
	key, err := promptSecretUntil(fmt.Sprintf("New key (at least %d characters): ", keystorage.MinPassphraseLength), check)
	if err != nil {
		return "", err
	}
	_, err = promptSecretUntil("Repeat the new key: ", func(repeat string) error {
		if repeat != key {
			return errors.New("that doesn't match")
		}
		return nil
	})
	return key, err
}
