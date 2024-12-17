package fuel

import (
	"checkers/account"
	"checkers/models"
	"checkers/modules/modulesHelpers"
	"checkers/utils"
	"fmt"
)

type Fuel struct {
	Endpoint string
}

func NewFuel() (*Fuel, error) {
	return &Fuel{
		Endpoint: "https://mainnet-14236c37.fuel.network/allocations?accounts=%s",
	}, nil
}

func (f *Fuel) Check(acc *account.Account) (float64, error) {
	client := modulesHelpers.CreateHttpClient(acc.Proxy)

	var fuelResp models.FuelResp
	if err := client.SendJSONRequest(fmt.Sprintf(f.Endpoint, acc.Address), "GET", nil, &fuelResp); err != nil {
		return 0, err
	}

	return utils.ConvertFrom9(fuelResp.Allocation[0].Amount), nil
}
