package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) GetTxInfo(c echo.Context) error {
	txHashStr := c.QueryParam("tx_hash")

	tx, receipt, err := d.service.GetTxInfo(txHashStr)

	txInfo := api.TxInfo{
		Tx:      tx,
		Receipt: receipt,
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(http.StatusOK, txInfo, "  ")
}
