package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	DBConfig DBConfig
}

type ServerConfig struct {
	Address string `envconfig:"SERVER_ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"SERVER_PORT" default:"8080"`
}

type DBConfig struct {
	Name     string `envconfig:"DB_DATABASE" default:"buddy"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASS" default:"pass"`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Host     string `envconfig:"DB_HOST" default:"db-buddy"`
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
