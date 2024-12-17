package account

type Account struct {
	Address string
	Module  string
	Proxy   string
}

func NewAccount(address, module, proxy string) (*Account, error) {
	return &Account{
		Address: address,
		Module:  module,
		Proxy:   proxy,
	}, nil
}
