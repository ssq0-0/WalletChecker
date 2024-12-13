package account

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	Address common.Address
	Module  string
	Proxy   string
}

func NewAccount(address, module string) (*Account, error) {
	return &Account{
		Address: common.HexToAddress(address),
		Module:  module,
	}, nil
}
