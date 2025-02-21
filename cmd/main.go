package main

import (
	"FinanceChecker/internal/config"
	sl "FinanceChecker/internal/lib/log"
	"FinanceChecker/internal/log"
	router "FinanceChecker/internal/ports"
	"net/http"
	"os"

	"FinanceChecker/internal/repo/sqlite"
)

func main() {
	// TODO: init config

	cfg := config.MustLoad()

	// TODO: init logger

	log := log.SetupLogger(cfg.Env)

	log.Info("app started")

	// TODO: init storage

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Error(err))
		os.Exit(1)
	}

	// TODO: init router
	err = http.ListenAndServe(":8050", router.SetupRouter(log, storage))
	if err != nil {
		log.Error("failed to start server")
	}
	// TODO: run server
}
