package helpers

import (
	"checkers/account"
	"checkers/config"
	"checkers/ethClient"
	"checkers/logger"
	"checkers/utils"
	"errors"
	"fmt"
)

func GetAllPath() (map[string]string, error) {
	accWalletsPath := "account/wallets.csv"
	configPath := "config/config.json"
	proxyPath := "account/proxy.txt"

	return map[string]string{
		"wallets": accWalletsPath,
		"config":  configPath,
		"proxy":   proxyPath,
	}, nil
}

func ClientsInit() (map[string]*ethClient.Client, error) {
	var clients = make(map[string]*ethClient.Client)
	for chain, rpc := range config.RPCs {
		client, err := ethClient.NewClient(rpc)
		if err != nil {
			logger.GlobalLogger.Errorf("Ошибка создания eth client для сети %s: %v", chain, err)
			continue
		}
		clients[chain] = client
	}

	if len(clients) == 0 {
		return nil, errors.New("не удалось создать ни одного клиента. Проверьте настройки RPC")
	}

	return clients, nil
}

func ProxyInit(proxyFilePath string) ([]string, error) {
	addresses, err := utils.FileReader(proxyFilePath)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func AccsInit(accAddressesPath, module string, proxys []string) ([]*account.Account, error) {
	addresses, err := utils.FileReader(accAddressesPath)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("нет ни одного адреса. Проверьте файл wallets.csv")
	}

	if len(proxys) == 0 {
		logger.GlobalLogger.Warnf("Прокси не обнаружено, возможны ошибки при проверке аллокации.")
		proxys = nil
	}

	var accs []*account.Account
	for i, addr := range addresses {
		var proxy string

		if len(proxys) > 0 {
			proxyIndex := i % len(proxys)
			parsedProxy, err := utils.ParseProxy(proxys[proxyIndex])
			if err != nil {
				return nil, fmt.Errorf("не удалось разобрать прокси: %v", err)
			}
			proxy = parsedProxy
		}

		acc, err := account.NewAccount(addr, module, proxy)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}
