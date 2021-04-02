package models

type Settings struct {
	OracleHost string `json:"oracle_host"`
	OraclePort string `json:"oracle_port"`
	OracleKey  string `json:"oracle_key"`
}

func (d Settings) GetOracleHost() string {
	return d.OracleHost
}

func (d Settings) GetOraclePort() string {
	return d.OraclePort
}

func (d Settings) GetOracleKey() string {
	return d.OracleKey
}
