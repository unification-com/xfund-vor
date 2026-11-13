package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
	"oracle/service"
)

// ScanFrom moves where the chain's scan reads next.
func (d *Oracle) ScanFrom(c echo.Context, svc *service.Service) error {
	var requestModel api.OracleScanFromRequestModel
	if err := decodeBody(c, &requestModel); err != nil {
		return badRequest(c, err)
	}
	if err := required("block", requestModel.Block); err != nil {
		return badRequest(c, err)
	}
	if svc.Scanner == nil {
		return c.String(http.StatusServiceUnavailable, ErrChainNotRunning.Error())
	}
	from, to, err := svc.Scanner.ScanFrom(requestModel.Block)
	if errors.Is(err, service.ErrInvalidScanTarget) {
		return badRequest(c, err)
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api.OracleScanFromResponseModel{From: from, To: to})
}
