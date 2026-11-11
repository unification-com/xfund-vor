package api

import (
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
)

func (d *Oracle) QueryFees(c echo.Context) error {
	var requestModel api.OracleQueryFeesModel
	// the body is optional: without one, the oracle's base fee is returned
	if err := decodeBody(c, &requestModel); err != nil && !errors.Is(err, io.EOF) {
		return badRequest(c, err)
	}
	if requestModel.Consumer != "" {
		if _, err := parseAddress("consumer", requestModel.Consumer); err != nil {
			return badRequest(c, err)
		}
	}
	fee, err := d.service.QueryFees(requestModel.Consumer)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fee)
}
