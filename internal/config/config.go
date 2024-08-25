package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server        ServerConfig
	DBConfig      DBConfig
	CognitoConfig CognitoConfig
	AWSConfig     AWSConfig
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

type CognitoConfig struct {
	ClientId   string `envconfig:"COGNITO_CLIENT_ID"`
	UserPoolId string `envconfig:"COGNITO_USER_POOL_ID"`
}

type AWSConfig struct {
	Region string `envconfig:"AWS_REGION"`
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
