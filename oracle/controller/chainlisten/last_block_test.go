package chainlisten

import "testing"

// A zero block number (e.g. from a misbehaving RPC) must be ignored: blockNumber - 1 used to wrap round to the
// maximum uint64 and be saved in the keystore as -1, breaking the next start.
func TestSetLastBlockNumberIgnoresZero(t *testing.T) {
	var listener VORCoordinatorListener // no service: saving anything would panic
	if err := listener.SetLastBlockNumber(0); err != nil {
		t.Fatalf("SetLastBlockNumber(0) = %v, want nil", err)
	}
}
