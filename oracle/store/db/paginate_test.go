package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"oracle/store/db"
)

// The requests handler divides by the limit: it must never be 0.
func TestPageAndLimit(t *testing.T) {
	for _, c := range []struct{ page, limit, wantPage, wantLimit int }{
		{0, 0, 1, 10},
		{-3, -1, 1, 10},
		{2, 20, 2, 20},
		{1, 1000, 1, 100},
	} {
		page, limit := db.PageAndLimit(c.page, c.limit)
		assert.Equal(t, c.wantPage, page)
		assert.Equal(t, c.wantLimit, limit)
	}
}
