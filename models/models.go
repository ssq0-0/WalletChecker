package models

type LineaResp struct {
	Address string `json:"user_address"`
	Xp      int    `json:"xp"`
}

type OdosResp struct {
	Data struct {
		PendingTokenBalance string `json:"pendingTokenBalance"`
	} `json:"data"`
}
