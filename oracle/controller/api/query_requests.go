package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
	"strconv"
)

func (d *Oracle) QueryRequests(c echo.Context) error {
	requestId, _  := strconv.Atoi(c.QueryParam("id"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	status, _ := strconv.Atoi(c.QueryParam("status"))
	order := c.QueryParam("order")

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	requests := &api.RequestResponse{}

	dbRequests, count, err := d.service.Requests(requestId, page, limit, status, order)

	numPages := count / int64(limit)
	if count % int64(limit) > 0 {
		numPages = numPages + 1
	}

	for _, reqRow := range dbRequests {
		res := api.RandomnessRequestModel{
			ID: reqRow.ID,
			CreatedAt: reqRow.CreatedAt,
			UpdatedAt: reqRow.UpdatedAt,
			Sender: reqRow.Sender,
			RequestId: reqRow.RequestId,
			RequestBlockNumber: reqRow.RequestBlockNumber,
			RequestTxHash: reqRow.RequestTxHash,
			RequestGasUsed: reqRow.RequestGasUsed,
			RequestGasPrice: reqRow.RequestGasPrice,
			Fee: reqRow.Fee,
			Randomness: reqRow.Randomness,
			FulfillBlockNumber: reqRow.FulfillBlockNumber,
			FulfillTxHash: reqRow.FulfillTxHash,
			FulfillGasUsed: reqRow.FulfillGasUsed,
			FulfillGasPrice: reqRow.FulfillGasPrice,
			Status: reqRow.Status,
			StatusText: reqRow.GetStatusString(),
			StatusReason: reqRow.StatusReason,
		}
		requests.Requests = append(requests.Requests, res)
	}

	requests.Pages.Page = uint(page)
	requests.Pages.NumPages = uint(numPages)
	requests.Pages.NumRecords = uint(count)
	requests.Pages.Limit = uint(limit)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, requests, "  ")
	}
	return c.JSONPretty(http.StatusOK, requests, "  ")
}
