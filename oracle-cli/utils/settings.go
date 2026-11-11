package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oraclecli/models"
	"os"
)

var Settings *SettingsStore

// settingsFileMode: the settings hold the oracle's API key, so only their owner may read them.
const settingsFileMode = 0600

type SettingsStore struct {
	path     string
	Settings *models.Settings
}

func (d *SettingsStore) String() string {
	return fmt.Sprintf(`
Current oracle-cli settings:
	Oracle Host: %s
	Oracle Port: %s
	Key: ...
`, d.Settings.GetOracleHost(), d.Settings.GetOraclePort())
}

// NewSettingsStore loads the settings at filePath, creating an empty settings file if there is none.
func NewSettingsStore(filePath string) (*SettingsStore, error) {
	store := &SettingsStore{path: filePath, Settings: &models.Settings{}}

	data, err := ioutil.ReadFile(filePath)
	if os.IsNotExist(err) {
		return store, store.save()
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, store.Settings); err != nil {
		return nil, err
	}
	// settings written by older versions are world-readable
	return store, os.Chmod(filePath, settingsFileMode)
}

func (d *SettingsStore) SetOracleHost(host string) (err error) {
	d.Settings.OracleHost = host
	err = d.save()
	return
}

func (d *SettingsStore) SetOraclePort(port string) (err error) {
	d.Settings.OraclePort = port
	err = d.save()
	return
}

func (d *SettingsStore) SetOracleKey(key string) (err error) {
	d.Settings.OracleKey = key
	err = d.save()
	return
}

func (d *SettingsStore) save() error {
	jsonByte, err := json.Marshal(d.Settings)
	if err != nil {
		return err
	}
	// os.WriteFile truncates, so a shorter rewrite can't leave a stale tail behind
	return os.WriteFile(d.path, jsonByte, settingsFileMode)
}
