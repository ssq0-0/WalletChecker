package pengu

import (
	"checkers/account"
	"checkers/logger"
)

type Pengu struct{}

func NewPengu() (*Pengu, error) {
	return &Pengu{}, nil
}

func (p *Pengu) Check(acc *account.Account) (float64, error) {
	logger.GlobalLogger.Infof("заглушка")
	return 0, nil
}
