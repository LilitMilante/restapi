package apiserver

import "github.com/LilitMilante/restapi/internal/app/store"

type Config struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
