package utils

import (
	"encoding/json"
	"io/ioutil"
	"oraclecli/models"
	"os"
)

var Settings *SettingsStore

type SettingsStore struct {
	File     *os.File
	Settings *models.Settings
}

func NewSettingsStore(filePath string) (*SettingsStore, error) {
	var err error
	var settingsFile *os.File
	var settings = models.Settings{}

	if _, err = os.Stat(filePath); err == nil {
		settingsFile, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(settingsFile)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &settings)
		if err != nil {
			return nil, err
		}
	} else if os.IsNotExist(err) {
		settingsFile, err = os.Create(filePath)
		_, err := settingsFile.Write([]byte(`{}`))
		return nil, err
	}

	return &SettingsStore{
		File:     settingsFile,
		Settings: &settings,
	}, err
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
	_, err = d.File.WriteAt(jsonByte, 0)
	return err
}
