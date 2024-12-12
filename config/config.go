package config

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Config struct {
	LineaCA string `json:"linea_ca"`
}

func init() {
	parsedABI, err := abi.JSON(bytes.NewReader(Erc20JSON))
	if err != nil {
		// logger.GlobalLogger.Fatalf("Ошибка при парсинге ABI: %v", err)
	}

	Erc20ABI = &parsedABI

}

func ConfigUpload(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
