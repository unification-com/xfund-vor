package api

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
	"oracle/service"
	"oracle/store/db"
	"strconv"
)

func (d *Oracle) QueryRequests(c echo.Context, svc *service.Service) error {
	requestId, _ := strconv.Atoi(c.QueryParam("id"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, limit = db.PageAndLimit(page, limit)
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		status = -1 // any status
	}
	order := c.QueryParam("order")

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	requests := &api.RequestResponse{}

	dbRequests, count, err := svc.Requests(requestId, page, limit, status, order)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	numPages := count / int64(limit)
	if count%int64(limit) > 0 {
		numPages = numPages + 1
	}

	for _, reqRow := range dbRequests {
		var seed uint64 // no seed until the request's fulfilment starts
		if seedBig, err := hexutil.DecodeBig(reqRow.Seed); err == nil {
			seed = seedBig.Uint64()
		}

		res := api.RandomnessRequestModel{
			ID:                 reqRow.ID,
			CreatedAt:          reqRow.CreatedAt,
			UpdatedAt:          reqRow.UpdatedAt,
			Sender:             reqRow.Sender,
			RequestId:          reqRow.RequestId,
			RequestBlockNumber: reqRow.RequestBlockNumber,
			RequestBlockHash:   reqRow.RequestBlockHash,
			RequestTxHash:      reqRow.RequestTxHash,
			RequestGasUsed:     reqRow.RequestGasUsed,
			RequestGasPrice:    reqRow.RequestGasPrice,
			SeedHex:            reqRow.Seed,
			Seed:               seed,
			Fee:                reqRow.Fee,
			Randomness:         reqRow.Randomness,
			FulfillBlockNumber: reqRow.FulfillBlockNumber,
			FulfillBlockHash:   reqRow.FulfillBlockHash,
			FulfillTxHash:      reqRow.FulfillTxHash,
			FulfillGasUsed:     reqRow.FulfillGasUsed,
			FulfillGasPrice:    reqRow.FulfillGasPrice,
			Status:             reqRow.Status,
			StatusText:         reqRow.GetStatusString(),
			StatusReason:       reqRow.StatusReason,
		}
		requests.Requests = append(requests.Requests, res)
	}

	requests.Pages.Page = uint(page)
	requests.Pages.NumPages = uint(numPages)
	requests.Pages.NumRecords = uint(count)
	requests.Pages.Limit = uint(limit)

	return c.JSONPretty(http.StatusOK, requests, "  ")
}
