package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/service"
)

type fakeScanner struct{ target string }

func (f *fakeScanner) ScanFrom(target string) (uint64, uint64, error) {
	f.target = target
	if target == "bad" {
		return 0, 0, fmt.Errorf("%q: %w", target, service.ErrInvalidScanTarget)
	}
	if target == "broken" {
		return 0, 0, errors.New("the database is down")
	}
	return 5, 900, nil
}

func TestScanFrom(t *testing.T) {
	scanner := &fakeScanner{}
	svc := &service.Service{Scanner: scanner}
	call := func(body string, svc *service.Service) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		ctx := echo.New().NewContext(httptest.NewRequest(http.MethodPost, "/scanfrom", strings.NewReader(body)), rec)
		require.NoError(t, (&Oracle{}).ScanFrom(ctx, svc))
		return rec
	}

	rec := call(`{"block":"head-100"}`, svc)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"from":5,"to":900}`, rec.Body.String())
	assert.Equal(t, "head-100", scanner.target)

	assert.Equal(t, http.StatusBadRequest, call(`{"block":"bad"}`, svc).Code, "an invalid target is the caller's mistake")
	assert.Equal(t, http.StatusBadRequest, call(`{}`, svc).Code)
	assert.Equal(t, http.StatusInternalServerError, call(`{"block":"broken"}`, svc).Code)
	assert.Equal(t, http.StatusServiceUnavailable, call(`{"block":"head"}`, &service.Service{}).Code, "no scanner yet")
}
