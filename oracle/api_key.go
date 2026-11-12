package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"oracle/config"
	"oracle/store/keystorage"
	"oracle/utils"
)

// newAPIKey replaces the HTTP API key, writing it to --out, or to api_key beside the config where the admin
// commands read it; the keystore keeps only its hash.
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
	out := options.APIKeyOut
	if out == "" {
		out = defaultAPIKeyFile()
	}
	if err = issueAPIKey(keystore, out); err != nil {
		return err
	}
	fmt.Println("New API key written to", out)
	return nil
}

// defaultAPIKeyFile is where the API key is kept for the admin commands: api_key beside the config.
func defaultAPIKeyFile() string {
	return filepath.Join(filepath.Dir(options.Config), "api_key")
}

// issueAPIKey makes a random API key, sets it, and writes it to out (0600, replacing any older key). The key
// is staged beside out first and only moved into place once set, so out never holds a key the oracle doesn't
// accept, and a failure leaves the old key working.
func issueAPIKey(keystore *keystorage.Keystorage, out string) error {
	key, err := keystorage.RandomKey()
	if err != nil {
		return err
	}
	staging := out + ".new"
	if err = os.WriteFile(staging, []byte(key+"\n"), 0600); err != nil {
		return fmt.Errorf("writing the API key: %w", err)
	}
	if err = keystore.SetAPIKey(key); err != nil {
		_ = os.Remove(staging)
		return err
	}
	if err = utils.RenameDurably(staging, out); err != nil {
		return fmt.Errorf("the new API key is set but is in %s, not %s (move it there): %w", staging, out, err)
	}
	return nil
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
