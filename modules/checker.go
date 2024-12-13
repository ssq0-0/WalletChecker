package modules

import "checkers/account"

type Checker interface {
	Check(acc *account.Account) (float64, error)
}
