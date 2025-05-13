package config

import "github.com/BurntSushi/toml"

type Config struct {
	Postgres struct {
		URL string `toml:"url"`
	} `toml:"postgres"`

	Server struct {
		Addr string `toml:"addr"`
		Mode string `toml:"mode"`
	} `toml:"server"`

	Log struct {
		Level string `toml:"level"`
		Path  string `toml:"path"`
	} `toml:"log"`
}

var Cfg = new(Config)

func Init(path string) (*Config, error) {
	_, err := toml.DecodeFile(path, Cfg)
	if err != nil {
		return nil, err
	}

	return Cfg, nil
}
