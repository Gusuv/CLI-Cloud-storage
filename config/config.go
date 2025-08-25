package config

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	User     string `config:"DB_USER"`
	Password string `config:"DB_PASSWORD"`
	Host     string `config:"DB_HOST"`
	Port     string `config:"DB_PORT"`
	Name     string `config:"DB_NAME"`
	SSLMode  string `config:"DB_SSL"`
}

func InitConfig() (*Config, error) {

	godotenv.Load(".env")

	loader := confita.NewLoader(env.NewBackend())

	cfg := &Config{}

	if err := loader.Load(context.Background(), cfg); err != nil {
		return nil, fmt.Errorf("load is failed: %w", err)
	}

	return cfg, nil
}
