package account

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	Address common.Address
}

func NewAccount(address string) (*Account, error) {
	return &Account{
		Address: common.HexToAddress(address),
	}, nil
}
