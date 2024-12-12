package modules

import (
	"checkers/config"
	"checkers/ethClient"
	"checkers/modules/linea"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"
)

func ModsInit(cfg *config.Config, clients map[string]*ethClient.Client) (map[string]Checker, error) {
	var (
		g        errgroup.Group
		checkers sync.Map
	)

	g.Go(func() error {
		linea, err := linea.NewLinea(common.HexToAddress(cfg.LineaCA), clients["linea"])
		if err != nil {
			return err
		}
		checkers.Store("linea", linea)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	result := make(map[string]Checker)
	checkers.Range(func(key, value any) bool {
		result[key.(string)] = value.(Checker)
		return true
	})

	return result, nil
}
