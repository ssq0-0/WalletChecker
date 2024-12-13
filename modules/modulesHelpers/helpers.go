package modulesHelpers

import (
	"checkers/httpClient"
	"checkers/logger"
)

func CreateHttpClient(proxy string) *httpClient.HttpClient {
	client, err := httpClient.NewHttpClient(&proxy)
	if err != nil {
		logger.GlobalLogger.Warnf("ошибка создания http клиента: %w", err)
		return nil
	}

	return client
}
