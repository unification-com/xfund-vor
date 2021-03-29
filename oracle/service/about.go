package service

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"oracle/config"
	"oracle/utils"
	"oracle/utils/walletworker"
)

func (d *Service) About() (response string, err error) {

	privateKey, err := crypto.HexToECDSA(utils.RemoveHexPrefix(d.Store.Keystorage.GetSelectedPrivateKey()))
	publicKey := hexutil.Encode(crypto.FromECDSAPub(privateKey.Public().(*ecdsa.PublicKey)))
	ECDSAoraclePublicKey, err := crypto.UnmarshalPubkey(crypto.FromECDSAPub(privateKey.Public().(*ecdsa.PublicKey)))
	_, oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)
	keyhash, _ := publicKey, d.VORCoordinatorCaller.HashOfKey()
	return fmt.Sprintf(`Account: %s
Private Key: %s
Public Key: %s
KeyHash: %s
Address: %s

Contract address (VORCoordinator): %s
Host: %s 
Port: %s
`, config.Conf.Keystorage.Account, d.Store.Keystorage.GetSelectedPrivateKey(), publicKey, common.BytesToHash([]byte(keyhash[:])), oracleAddress, config.Conf.VORCoordinatorContractAddress, config.Conf.Serve.Host, string(config.Conf.Serve.Port)), nil
}
