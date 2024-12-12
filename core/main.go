package main

import (
	"checkers/config"
	"checkers/core/helpers"
	"checkers/modules"
)

func main() {
	paths, err := helpers.GetAllPath()
	if err != nil {
		return
	}

	accs, err := helpers.AccsInit(paths["wallets"])
	if err != nil {
		return
	}

	clients, err := helpers.ClientsInit()
	if err != nil {
		return
	}

	config, err := config.ConfigUpload(paths["config"])
	if err != nil {
		return
	}

	mods, err := modules.ModsInit(config, clients)
	if err != nil {
		return
	}

	for _, acc := range accs {
		mods["linea"].Check(acc)
	}
}
