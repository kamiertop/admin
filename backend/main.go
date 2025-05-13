package main

import (
	"backend/config"
	"backend/storage"
	"flag"
)

func init() {
	configPath := flag.String("config", "config.toml", "config file path, default is .")
	flag.Parse()

	cfg, err := config.Init(*configPath)
	if err != nil {
		panic(err)
	}
	if err := storage.InitPostgres(cfg.Postgres.URL); err != nil {
		panic(err)
	}
}

func main() {

}
