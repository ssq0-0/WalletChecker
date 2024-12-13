package main

import (
	"checkers/account"
	"checkers/analisor"
	"checkers/config"
	"checkers/core/helpers"
	"checkers/logger"
	"checkers/modules"
	"regexp"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	paths, err := helpers.GetAllPath()
	if err != nil {
		logger.GlobalLogger.Error(err)
		return
	}

	clients, err := helpers.ClientsInit()
	if err != nil {
		logger.GlobalLogger.Error(err)
		return
	}

	config, err := config.ConfigUpload(paths["config"])
	if err != nil {
		logger.GlobalLogger.Error(err)
		return
	}

	accs, err := helpers.AccsInit(paths["wallets"], userChoice())
	if err != nil {
		logger.GlobalLogger.Error(err)
		return
	}

	mods, err := modules.ModsInit(config, clients)
	if err != nil {
		return
	}
	aggregator := analisor.NewAggreagtor()
	proccessAccount(accs, mods, aggregator)
}

func userChoice() string {
	modules := []string{
		"1. Linea LXP",
		"0. Выйти.",
	}

	var selected string
	if err := survey.AskOne(&survey.Select{
		Message: "Выберите чекер.",
		Options: modules,
		Default: modules[len(modules)-1],
	}, &selected); err != nil {
		logger.GlobalLogger.Errorf("Ошибка выбора модуля: %v", err)
		return ""
	}

	rgx := regexp.MustCompile(`^\d+\.\s*`)
	selected = rgx.ReplaceAllString(selected, "")
	return selected
}

func proccessAccount(accs []*account.Account, mods map[string]modules.Checker, aggregator *analisor.Aggregator) {
	var (
		wg      sync.WaitGroup
		workers = 10
		sem     = make(chan struct{}, workers)
	)
	for _, acc := range accs {
		wg.Add(1)
		sem <- struct{}{}

		go func(acc *account.Account) {
			defer wg.Done()
			defer func() { <-sem }()

			res, err := mods[acc.Module].Check(acc)
			if err != nil {
				return
			}
			aggregator.Add(acc.Address, res)
		}(acc)
	}
	wg.Wait()

	aggregator.LogAnalizedData()
	if err := aggregator.WriteAnalizedData(); err != nil {
		logger.GlobalLogger.Error(err)
	}
}