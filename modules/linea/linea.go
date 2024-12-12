package linea

import (
	"checkers/account"
	"checkers/ethClient"
	"checkers/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Linea struct {
	LxpCA  common.Address
	Client *ethClient.Client
}

func NewLinea(lxpCA common.Address, ethClient *ethClient.Client) (*Linea, error) {
	return &Linea{
		LxpCA:  lxpCA,
		Client: ethClient,
	}, nil
}

func (l *Linea) Check(acc *account.Account) float64 {
	result, err := l.Client.BalanceCheck(acc.Address, l.LxpCA)
	if err != nil {
		return 0
	}

	return utils.ConvertFrom18(result)
}
