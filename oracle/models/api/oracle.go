package api

import "time"

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

type OracleChangeGranularFeeRequestModel struct {
	Consumer string `json:"consumer"`
	Amount   int64  `json:"amount"`
}

type OracleQueryFeesModel struct {
	Consumer string `json:"consumer"`
}

type RandomnessRequestModel struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"created"`
	UpdatedAt          time.Time `json:"updated"`
	Sender             string    `json:"consumer"`
	RequestId          string    `json:"request_id"`
	RequestBlockNumber uint64    `json:"request_block_num"`
	RequestTxHash      string    `json:"request_tx_hash"`
	RequestGasUsed     uint64    `json:"request_gas"`
	RequestGasPrice    uint64    `json:"request_gas_price"`
	Fee                uint64    `json:"fee"`
	Randomness         string    `json:"randomness"`
	FulfillBlockNumber uint64    `json:"fulfill_block_num"`
	FulfillTxHash      string    `json:"fulfill_tx_hash"`
	FulfillGasUsed     uint64    `json:"fulfill_gas"`
	FulfillGasPrice    uint64    `json:"fulfill_gas_price"`
	Status             int       `json:"status"`
	StatusText         string    `json:"status_text"`
	StatusReason       string    `json:"status_reason"`
}

type Pages struct {
	Page       uint `json:"page"`
	NumPages   uint `json:"num_pages"`
	NumRecords uint `json:"num_records"`
	Limit      uint `json:"limit"`
}
type RequestResponse struct {
	Requests []RandomnessRequestModel `json:"requests"`
	Pages    Pages `json:"pagination"`
}

type AnalyticsResponse struct {
	GasUsedMax           uint64
	GasUsedMin           uint64
	GasUsedMean          uint64
	GasPriceMax          uint64
	GasPriceMin          uint64
	GasPriceMean         uint64
	CostMaxEth           float64
	CostMinEth           float64
	CostMeanEth          float64
	TotalCostEth         float64
	TotalFeesEarnedXfund float64
	TotalFeesEarnedEth   float64
	ProfitLossEth        float64
	NumberAnalysed       uint64
}
