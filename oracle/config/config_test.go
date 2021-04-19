package config_test

import (
	"oracle/config"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	configuration, err := config.NewConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(*configuration)
}
