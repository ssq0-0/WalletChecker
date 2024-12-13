package linea

import (
	"checkers/account"
	"checkers/ethClient"
	"checkers/httpClient"
	"checkers/logger"
	"checkers/models"
	"checkers/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Linea struct {
	LxpCA    common.Address
	Endpoint string
	Client   *ethClient.Client
}

func NewLinea(lxpCA common.Address, ethClient *ethClient.Client) (*Linea, error) {
	return &Linea{
		LxpCA:    lxpCA,
		Client:   ethClient,
		Endpoint: "https://kx58j6x5me.execute-api.us-east-1.amazonaws.com/linea/getUserPointsSearch?user=",
	}, nil
}

func (l *Linea) Check(acc *account.Account) (float64, error) {
	switch acc.Module {
	case "Linea LXP":
		result, err := l.Client.BalanceCheck(acc.Address, l.LxpCA)
		if err != nil {
			return 0, err
		}

		return utils.ConvertFrom18(result), nil
	case "Linea LXP-l":
		client := createHttpClient(acc.Proxy)
		url := l.createRequestURL(acc)

		var lineaResp []models.LineaResp
		if err := client.SendJSONRequest(url, "GET", nil, &lineaResp); err != nil {
			return 0, err
		}

		return float64(lineaResp[0].Xp), nil
	}
	return 0, nil
}

func (l *Linea) createRequestURL(acc *account.Account) string {
	return l.Endpoint + strings.ToLower(acc.Address.Hex())
}

func createHttpClient(proxy string) *httpClient.HttpClient {
	client, err := httpClient.NewHttpClient(&proxy)
	if err != nil {
		logger.GlobalLogger.Warnf("ошибка создания http клиента: %w", err)
		return nil
	}

	return client
}
