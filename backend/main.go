package main

import (
	"flag"

	"backend/api/router"
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
	if err := db.InitPostgres(cfg.Postgres.URL); err != nil {
		panic(err)
	}

	if err := router.Serve(cfg.Server.Addr); err != nil {
		panic(err)
	}
}
