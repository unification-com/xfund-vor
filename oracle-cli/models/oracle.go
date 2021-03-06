package models

type OracleWithdrawRequestModel struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type OracleChangeFeeRequestModel struct {
	Amount int64 `json:"amount"`
}

type OracleChangeGranularFeeRequestModel struct {
	Consumer string `json:"consumer"`
	Amount   int64  `json:"amount"`
}

type OracleRegisterRequestModel struct {
	AccountName string `json:"account_name"`
	PrivateKey  string `json:"private_key"`
	Fee         int64  `json:"fee"`
}

type OracleQueryFeesModel struct {
	Consumer string `json:"consumer"`
}
