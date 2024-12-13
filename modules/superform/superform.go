package superform

import (
	"checkers/account"
	"checkers/models"
	"checkers/modules/modulesHelpers"
	"fmt"
	"log"
)

type Superform struct {
	CredEndpoint string
}

func NewSuperform() (*Superform, error) {
	return &Superform{
		CredEndpoint: "https://www.superform.xyz/api/proxy/superrewards/exploration/leaderboard/cred/%s",
	}, nil
}

func (s *Superform) Check(acc *account.Account) (float64, error) {
	client := modulesHelpers.CreateHttpClient(acc.Proxy)

	var sfResp models.SuperformResp
	if err := client.SendJSONRequest(fmt.Sprintf(s.CredEndpoint, acc.Address), "GET", nil, &sfResp); err != nil {
		return 0, err
	}

	log.Println(sfResp.CurrentUser.Cred)
	return 0, nil
}
