package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"oracle/tools/secp256k1"
	"oracle/tools/vor"
	"oracle/utils"
)

func (d *Service) FulfillRandomness(seed vor.Seed, blockHash common.Hash, blockNum int64) (tx *types.Transaction, err error) {
	preSeed := vor.PreSeedData{
		PreSeed:   seed,
		BlockHash: blockHash,
		BlockNum:  uint64(blockNum),
	}
	oraclePrivateKeyECDSA, err := crypto.HexToECDSA(utils.RemoveHexPrefix(d.Store.Keystorage.GetSelectedPrivateKey()))

	secretKeyScalar := secp256k1.IntToScalar(oraclePrivateKeyECDSA.D)
	secretKey := secp256k1.ScalarToHash(secretKeyScalar)

	//secretKey, err := d.VORCoordinatorCaller.HashOfKey()
	if err != nil {
		return nil, err
	}

	marshalledResponse, err := vor.GenerateProofResponse(secretKey, preSeed)
	if err != nil {
		return nil, err
	}

	tx, err = d.VORCoordinatorCaller.FulfillRandomnessRequest(marshalledResponse[:])

	return
}
