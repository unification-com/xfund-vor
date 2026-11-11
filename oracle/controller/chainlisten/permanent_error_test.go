package chainlisten

import (
	"errors"
	"testing"
)

// Covers the fee-revert runaway guard: a permanent revert (e.g. a request whose paid fee is below the
// registered fee) must be classified permanent so it is failed terminally instead of re-sent forever;
// genuinely transient errors must NOT be, so they keep their normal capped retry.
func TestIsPermanentFulfilmentError(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"below agreed payment (the fee runaway)", errors.New("execution reverted: Below agreed payment"), true},
		{"execution reverted bare", errors.New("execution reverted"), true},
		{"legacy always-failing transaction", errors.New("gas required exceeds allowance, or always failing transaction"), true},
		{"case-insensitive", errors.New("EXECUTION REVERTED"), true},
		{"nonce too low (transient)", errors.New("nonce too low"), false},
		{"rpc blip (transient)", errors.New("connection refused"), false},
		{"replacement underpriced (transient)", errors.New("replacement transaction underpriced"), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isPermanentFulfilmentError(c.err); got != c.want {
				t.Fatalf("isPermanentFulfilmentError(%v) = %v, want %v", c.err, got, c.want)
			}
		})
	}
}
