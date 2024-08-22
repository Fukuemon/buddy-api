package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Address string `envconfig:"SERVER_ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"SERVER_PORT" default:"8080"`
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})
	return &config
}
