package models

import "math/big"

type LineaResp struct {
	Address string `json:"user_address"`
	Xp      int    `json:"xp"`
}

type OdosResp struct {
	Data struct {
		PendingTokenBalance string `json:"pendingTokenBalance"`
	} `json:"data"`
}

type SuperformResp struct {
	CurrentUser struct {
		Cred float64 `json:"cred"`
	} `json:"current_user"`
}

type PenguResp struct {
	Total int `json:"total"`
}

type FuelResp struct {
	Allocation []struct {
		Amount *big.Int `json:"amount"`
	} `json:"allocations"`
}
