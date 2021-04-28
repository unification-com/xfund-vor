package service

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
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

	tokens, err := d.VORCoordinatorCaller.QueryWithdrawableTokens()
	var withdrawableTokens = ""
	if err != nil {
		withdrawableTokens = err.Error()
	} else {
		toXfund := new(big.Float).Quo(new(big.Float).SetInt(tokens), big.NewFloat(params.GWei))
		withdrawableTokens = fmt.Sprintf("%s (%s XFUND)", tokens.String(), toXfund.String())
	}

	balance, err := d.VORCoordinatorCaller.GetOracleEthBalance()
    var ethBalance  = ""
	if err != nil {
		ethBalance = err.Error()
	} else {
		toEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(params.Ether))
		ethBalance = fmt.Sprintf("%s (%s ETH)", balance.String(), toEth.String())
	}

	return fmt.Sprintf(`
VORCoordinator address: %s
Host:                   %s 
Port:                   %d
Network:                %d

Account:                %s
Public Key:             %s
KeyHash:                %s
Address:                %s

Withdrawable Tokens:    %s
ETH Balance:            %s
`,
config.Conf.VORCoordinatorContractAddress,
config.Conf.Serve.Host,
config.Conf.Serve.Port,
config.Conf.NetworkID,
config.Conf.Keystorage.Account,
publicKey,
common.BytesToHash([]byte(keyhash[:])),
oracleAddress, withdrawableTokens,
ethBalance), nil
}
