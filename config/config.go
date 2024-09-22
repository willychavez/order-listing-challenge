package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Environment string `env:"ENV,required"`

	DB struct {
		Driver   string `env:"DB_DRIVER,required"`
		Host     string `env:"DB_HOST,required"`
		Port     string `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required"`
		Password string `env:"DB_PASSWORD,required"`
		Name     string `env:"DB_NAME,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	GraphQL struct {
		Port string `env:"GRAPHQL_PORT,required"`
	}
}

func Get() Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(fmt.Sprintf("Failed to decode: %s", err))
	}
	return cfg
}
