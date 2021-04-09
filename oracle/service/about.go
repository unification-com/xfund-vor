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
	keyhash, err := d.VORCoordinatorCaller.HashOfKey()

	return fmt.Sprintf(`
Contract address (VORCoordinator): %s
Host: %s 
Port: %d
Network: %d

Account: %s
Private Key: %s
Public Key: %s
KeyHash: %s
Address: %s

`, config.Conf.VORCoordinatorContractAddress, config.Conf.Serve.Host, config.Conf.Serve.Port, config.Conf.NetworkID, config.Conf.Keystorage.Account, d.Store.Keystorage.GetSelectedPrivateKey(), publicKey, common.BytesToHash([]byte(keyhash[:])), oracleAddress), nil
}
