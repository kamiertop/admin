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

	if duckDB, err := db.InitDuckDB(cfg.DuckDB.FilePath); err != nil {
		panic(err)
	} else {
		_ = duckDB
	}

	logger.Info("connect duckdb success")

	if err := router.Serve(cfg.Server.Addr, logger); err != nil {
		panic(err)
	}
}
