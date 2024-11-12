package config

import (
	"strings"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server        ServerConfig
	DBConfig      DBConfig
	CognitoConfig CognitoConfig
	AWSConfig     AWSConfig
}

type ServerConfig struct {
	Address          string        `envconfig:"SERVER_ADDRESS" default:"0.0.0.0"`
	Port             string        `envconfig:"SERVER_PORT" default:"8080"`
	AllowOrigins     []string      `envconfig:"ALLOW_ORIGINS" default:"http://localhost:3000"`
	AllowMethods     []string      `envconfig:"ALLOW_METHODS" default:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowHeaders     []string      `envconfig:"ALLOW_HEADERS" default:"Content-Length,Content-Type,Authorization"`
	AllowCredentials bool          `envconfig:"ALLOW_CREDENTIALS" default:"false"`
	MaxAge           time.Duration `envconfig:"MAX_AGE" default:"12h"`
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

func parseCommaSeparatedValues(input []string) []string {
	if len(input) == 1 {
		return strings.Split(input[0], ",")
	}
	return input
}

func (s *ServerConfig) parseAllowOrigins() {
	s.AllowOrigins = parseCommaSeparatedValues(s.AllowOrigins)
	s.AllowMethods = parseCommaSeparatedValues(s.AllowMethods)
	s.AllowHeaders = parseCommaSeparatedValues(s.AllowHeaders)
}

func GetConfig() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})
	config.Server.parseAllowOrigins()
	return &config
}
