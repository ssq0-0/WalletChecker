package pengu

import (
	"checkers/account"
	"checkers/models"
	"checkers/modules/modulesHelpers"
	"fmt"
)

type Pengu struct {
	Endpoint string
}

func NewPengu() (*Pengu, error) {
	return &Pengu{
		Endpoint: "https://api.clusters.xyz/v0.1/airdrops/pengu/eligibility/%s?",
	}, nil
}

func (p *Pengu) Check(acc *account.Account) (float64, error) {
	client := modulesHelpers.CreateHttpClient(acc.Proxy)

	var result models.PenguResp
	if err := client.SendJSONRequest(fmt.Sprintf(p.Endpoint, acc.Address), "GET", nil, &result); err != nil {
		return 0, err
	}

	return float64(result.Total), nil
}
