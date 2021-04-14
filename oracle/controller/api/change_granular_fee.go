package api

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
	"regexp"
)

func (d *Oracle) ChangeGranularFee(c echo.Context) error {
	var requestModel api.OracleChangeGranularFeeRequestModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(requestModel.Consumer) {
		return c.String(http.StatusInternalServerError, "not an address")
	}

	if !common.IsHexAddress(requestModel.Consumer) {
		return c.String(http.StatusInternalServerError, "not a hex address")
	}

	address := common.HexToAddress(requestModel.Consumer)

	transactionInfo, err := d.service.ChangeGranularFee(address, requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
