package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"oracle/tools/vor"
)

func (d *Oracle) FulfillRandomness(seed vor.Seed, blockHash common.Hash, blockNum int64) (tx *types.Transaction, err error) {
	preSeed := vor.PreSeedData{
		PreSeed:   seed,
		BlockHash: blockHash,
		BlockNum:  uint64(blockNum),
	}
	secretKey, err := d.VORCoordinatorCaller.HashOfKey()
	if err != nil {
		return nil, err
	}
	fmt.Println(secretKey)
	marshalledResponse, err := vor.GenerateProofResponse(secretKey, preSeed)
	if err != nil {
		return nil, err
	}
	fmt.Println(marshalledResponse)
	tx, err = d.VORCoordinatorCaller.FulfillRandomnessRequest(marshalledResponse[:])
	fmt.Println(tx)
	return
}
