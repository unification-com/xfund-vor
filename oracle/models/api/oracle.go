package api

type OracleWithdrawRequestModel struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type OracleRegisterRequestModel struct {
	AccountName     string `json:"account_name"`
	PrivateKey      string `json:"private_key"`
	Fee             int64  `json:"fee"`
}

type OracleChangeFeeRequestModel struct {
	Amount int64 `json:"amount"`
}
