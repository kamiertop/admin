package main

import (
	"flag"

	"backend/config"
	"backend/router"
	"backend/storage"
)

func main() {
	configPath := flag.String("config", "config.toml", "config file path, default is .")
	flag.Parse()

	cfg, err := config.Init(*configPath)
	if err != nil {
		panic(err)
	}
	if err := storage.InitPostgres(cfg.Postgres.URL); err != nil {
		panic(err)
	}

	if err := router.Serve(cfg.Server.Addr); err != nil {
		panic(err)
	}
}
