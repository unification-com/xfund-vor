package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

// The admin handlers check their input before anything reaches the chain: an undecodable body used to leave
// zero values (a withdrawal to 0x0, a zero fee) that went on to be sent.

// decodeBody decodes the request's JSON body into v.
func decodeBody(c echo.Context, v interface{}) error {
	if err := json.NewDecoder(c.Request().Body).Decode(v); err != nil {
		return fmt.Errorf("the request body is not valid JSON: %w", err)
	}
	return nil
}

// badRequest answers 400 with err's message.
func badRequest(c echo.Context, err error) error {
	return c.String(http.StatusBadRequest, err.Error())
}

// parseAddress parses a 0x-prefixed, non-zero address.
func parseAddress(name, hex string) (common.Address, error) {
	if len(hex) != 42 || !common.IsHexAddress(hex) {
		return common.Address{}, fmt.Errorf("%s is not a 0x-prefixed address", name)
	}
	address := common.HexToAddress(hex)
	if address == (common.Address{}) {
		return common.Address{}, fmt.Errorf("%s must not be the zero address", name)
	}
	return address, nil
}

// positive checks an amount or fee is above zero.
func positive(name string, n int64) error {
	if n <= 0 {
		return fmt.Errorf("%s must be above zero", name)
	}
	return nil
}

// required checks a field is present.
func required(name, value string) error {
	if value == "" {
		return errors.New(name + " is required")
	}
	return nil
}
