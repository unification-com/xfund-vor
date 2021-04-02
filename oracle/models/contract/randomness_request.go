package contract

import (
	"math/big"
)

type LogRandomnessRequest struct {
	KeyHash      []byte
	ConsumerSeed *big.Int
	FeePaid      *big.Int
}
