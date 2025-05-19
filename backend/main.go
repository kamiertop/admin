package main

import (
	"flag"

	"backend/api/router"
	"backend/common/log"
	"backend/config"
	"backend/dal/db"
)

func main() {
	configPath := flag.String("config", "config.toml", "config file path, default is .")
	flag.Parse()

	cfg, err := config.Init(*configPath)
	if err != nil {
		panic(err)
	}

	logger := log.Init(cfg.Log)

	if err := db.InitPostgres(cfg.Postgres.URL); err != nil {
		panic(err)
	}

	logger.Info("connect postgres success")

	if err := router.Serve(cfg.Server.Addr); err != nil {
		panic(err)
	}
}
