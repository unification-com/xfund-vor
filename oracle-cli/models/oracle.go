package models

type OracleWithdrawRequestModel struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}
