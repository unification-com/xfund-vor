package chainlisten

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"oracle/service"
)

// Range-limited RPCs reject wide eth_getLogs ranges (QL1's caps at 4,000 blocks), so a backlog is scanned in
// batches, and one poll's catch-up is bounded.
func TestScanRanges(t *testing.T) {
	assert.Empty(t, scanRanges(101, 100, 2000, 100), "nothing new since the last scan")
	assert.Equal(t, []blockRange{{100, 100}}, scanRanges(100, 100, 2000, 100))
	assert.Equal(t, []blockRange{{1, 2000}, {2001, 4000}, {4001, 4500}}, scanRanges(1, 4500, 2000, 100))
	assert.Equal(t, []blockRange{{1, 10}, {11, 20}}, scanRanges(1, 1000, 10, 2), "at most maxBatches per poll")
}

// A lagging RPC node can report a head behind a block it has already served: that must not wrap round to a
// huge age (which would fail the request as too old).
func TestBlocksSince(t *testing.T) {
	assert.Equal(t, uint64(5), blocksSince(105, 100))
	assert.Equal(t, uint64(0), blocksSince(100, 100))
	assert.Equal(t, uint64(0), blocksSince(99, 100))
}

// A scan position is a block, "head" or "head-N", from block 1 to the block after the head.
func TestScanTarget(t *testing.T) {
	const head = 1000
	for target, want := range map[string]uint64{"head": 1000, "head-100": 900, "head-999": 1, "1": 1, "500": 500, "1001": 1001} {
		got, err := scanTarget(target, head)
		assert.NoError(t, err, target)
		assert.Equal(t, want, got, target)
	}
	for _, target := range []string{"", "0", "1002", "head-1000", "head-x", "head+1", "-5", "latest"} {
		_, err := scanTarget(target, head)
		assert.ErrorIs(t, err, service.ErrInvalidScanTarget, target)
	}
}
