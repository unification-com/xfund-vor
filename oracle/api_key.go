package main

import (
	"fmt"
	"net"
	"os"

	"oracle/config"
	"oracle/store/keystorage"
)

// newAPIKey replaces the HTTP API key (what oracle-cli settings calls the oracle key). The new key is written
// to --out, or printed when there is none; the keystore keeps only its hash.
//
//	oracle api-key -c config.json [--out FILE]
func newAPIKey() error {
	if err := ensureOracleStopped(); err != nil {
		return err
	}
	keystore, err := keystorage.NewKeyStorage(log, config.Conf.Keystorage.File)
	if err != nil {
		return err
	}
	key, err := issueAPIKey(keystore, options.APIKeyOut)
	if err != nil {
		return err
	}
	if options.APIKeyOut != "" {
		fmt.Println("New API key written to", options.APIKeyOut)
		return nil
	}
	fmt.Println("New API key (put it in oracle-cli settings; it isn't shown again):")
	fmt.Println(key)
	return nil
}

// issueAPIKey makes a random API key and sets it, writing it to out first when out is given (a new file, at
// 0600), so a key is never set that nobody has a copy of.
func issueAPIKey(keystore *keystorage.Keystorage, out string) (string, error) {
	key, err := keystorage.RandomKey()
	if err != nil {
		return "", err
	}
	if out != "" {
		f, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			return "", fmt.Errorf("writing the API key: %w", err)
		}
		_, err = fmt.Fprintln(f, key)
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		if err != nil {
			_ = os.Remove(out)
			return "", fmt.Errorf("writing the API key: %w", err)
		}
	}
	if err = keystore.SetAPIKey(key); err != nil {
		if out != "" {
			_ = os.Remove(out)
		}
		return "", err
	}
	return key, nil
}

// ensureOracleStopped refuses to rewrite the keystore while the oracle runs: it holds the keystore in memory
// and saves it every poll, which would undo the change. A running oracle (of any version) is recognised by its
// API port being taken.
func ensureOracleStopped() error {
	address := fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("something is listening on %s, probably the oracle: stop it first (%v)", address, err)
	}
	return listener.Close()
}
