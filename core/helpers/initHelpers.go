package helpers

import (
	"checkers/account"
	"checkers/config"
	"checkers/ethClient"
	"checkers/logger"
	"checkers/utils"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GetAllPath() (map[string]string, error) {
	root := utils.GetRootDir()

	accWalletsPath := filepath.Join(root, "account", "wallets.csv")
	if _, err := os.Stat(accWalletsPath); os.IsNotExist(err) {
		return nil, err
	}

	configPath := filepath.Join(root, "config", "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, err
	}

	proxyPath := filepath.Join(root, "account", "proxy.txt")
	if _, err := os.Stat(proxyPath); os.IsNotExist(err) {
		return nil, err
	}

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
		proxys = nil
	}
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	var accs []*account.Account
	for i, addr := range addresses {
		var proxy string

		if len(proxys) > 0 {
			if len(proxys) < len(addresses) {
				proxy = proxys[randGen.Intn(len(proxys))]
			} else if len(proxys) == len(addresses) {
				var err error
				proxy, err = utils.ParseProxy(proxys[i])
				if err != nil {
					return nil, fmt.Errorf("не удалось разобрать прокси: %v", err)
				}
			}
		}

		acc, err := account.NewAccount(addr, module, proxy)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}

	return accs, nil

}
