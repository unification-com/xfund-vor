package api

import (
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

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
	RequestBlockHash   string    `json:"request_block_hash"`
	RequestTxHash      string    `json:"request_tx_hash"`
	RequestGasUsed     uint64    `json:"request_gas"`
	RequestGasPrice    uint64    `json:"request_gas_price"`
	SeedHex            string    `json:"seed_hex"`
	Seed               uint64    `json:"seed"`
	Fee                uint64    `json:"fee"`
	Randomness         string    `json:"randomness"`
	FulfillBlockNumber uint64    `json:"fulfill_block_num"`
	FulfillBlockHash   string    `json:"fulfill_block_hash"`
	FulfillTxHash      string    `json:"fulfill_tx_hash"`
	FulfillGasUsed     uint64    `json:"fulfill_gas"`
	FulfillGasPrice    uint64    `json:"fulfill_gas_price"`
	Status             int       `json:"status"`
	StatusText         string    `json:"status_text"`
	StatusReason       string    `json:"status_reason"`
}

type TxInfo struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
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
