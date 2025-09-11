package main

import (
	"flag"

	"backend/api/router"
	"backend/common/log"
	"backend/config"
	"backend/dal/db"
)

func main() {
	configPath := flag.String("config", "config.toml", "config file path, default is ./config.toml")
	flag.Parse()

	cfg, err := config.Init(*configPath)
	if err != nil {
		panic(err)
	}

	logger := log.Init(cfg.Log)

	if _, err := db.InitSqlite(cfg.Sqlite.FilePath, cfg.Sqlite.MaxIdleConn, cfg.Sqlite.MaxOpenConn); err != nil {
		panic(err)
	}
	logger.Info("connect sqlite3 success")

	if err := router.Serve(cfg.Server.Addr, logger); err != nil {
		panic(err)
	}
}
