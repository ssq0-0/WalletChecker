package ethClient

import (
	"checkers/config"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client *ethclient.Client
}

func NewClient(rpc string) (*Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
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
