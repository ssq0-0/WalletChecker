package helpers

import (
	"checkers/account"
	"checkers/config"
	"checkers/ethClient"
	"checkers/logger"
	"checkers/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

	return map[string]string{
		"wallets": accWalletsPath,
		"config":  configPath,
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

func AccsInit(accAddressesPath, module string) ([]*account.Account, error) {
	addresses, err := utils.FileReader(accAddressesPath)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("нет ни одного адреса. Проверьте файл wallets.csv")
	}

	var accs []*account.Account
	for _, addr := range addresses {
		acc, err := account.NewAccount(addr, module)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}
