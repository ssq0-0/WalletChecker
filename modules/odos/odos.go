package odos

import (
	"checkers/account"
	"checkers/models"
	"checkers/modules/modulesHelpers"
	"checkers/utils"
	"fmt"
	"strings"
)

type Odos struct {
	Endpoint string
}

func NewOdos() (*Odos, error) {
	return &Odos{
		Endpoint: "https://api.odos.xyz/loyalty/users/%s/balances",
	}, nil
}

func (o *Odos) Check(acc *account.Account) (float64, error) {
	client := modulesHelpers.CreateHttpClient(acc.Proxy)
	url := o.createRequestURL(acc)

	var odosResp models.OdosResp
	if err := client.SendJSONRequest(url, "GET", nil, &odosResp); err != nil {
		return 0, err
	}
	return utils.ConvertFrom18String(odosResp.Data.PendingTokenBalance), nil
}

func (o *Odos) createRequestURL(acc *account.Account) string {
	return fmt.Sprintf(o.Endpoint, strings.ToLower(acc.Address.Hex()))
}
