package ethClient

import (
	"checkers/config"
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client   *ethclient.Client
	FilePath string
	Txs      sync.Map
}

func NewClient(rpc string, filepath string) (*Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:   client,
		FilePath: filepath,
	}, nil
}

func CloseAllClients(clients map[string]*Client) {
	for _, client := range clients {
		if client.Client != nil {
			client.Client.Close()
		}
	}
}

func (c *Client) BalanceCheck(owner, tokenAddr common.Address) (*big.Int, error) {
	// if utils.IsNativeToken(tokenAddr) {
	// 	balance, err := c.Client.BalanceAt(context.Background(), owner, nil)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to get native coin balance: %v", err)
	// 	}
	// 	return balance, nil
	// }

	data, err := config.Erc20ABI.Pack("balanceOf", owner)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %v", err)
	}

	result, err := c.CallCA(tokenAddr, data)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %v", err)
	}

	var balance *big.Int
	if err := config.Erc20ABI.UnpackIntoInterface(&balance, "balanceOf", result); err != nil {
		return nil, fmt.Errorf("failed to unpack result: %v", err)
	}

	return balance, nil
}

func (c *Client) CallCA(toCA common.Address, data []byte) ([]byte, error) {
	callMsg := ethereum.CallMsg{
		To:   &toCA,
		Data: data,
	}

	return c.Client.CallContract(context.Background(), callMsg, nil)
}

func (c *Client) GetGasValues(msg ethereum.CallMsg) (uint64, *big.Int, *big.Int, error) {
	header, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil, nil, err
	}
	baseFee := header.BaseFee

	maxPriorityFeePerGas := big.NewInt(1e7)
	maxFeePerGas := new(big.Int).Add(baseFee, maxPriorityFeePerGas)

	gasLimit, err := c.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, nil, nil, err
	}

	return gasLimit, maxPriorityFeePerGas, maxFeePerGas, nil
}

func (c *Client) GetNonce(address common.Address) uint64 {
	nonce, err := c.Client.PendingNonceAt(context.Background(), address)
	if err != nil {
		// logger.GlobalLogger.Warnf("Failed to get nonce for address %s: %v", address.Hex(), err)
		return 0
	}
	return nonce
}
