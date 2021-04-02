package config_test

import (
	"oracle/config"
	"testing"
)

func TestNewConfig(t *testing.T) {
	configuration, err := config.NewConfig("../config.json")
	t.Error(err)
	t.Log(*configuration)
}
