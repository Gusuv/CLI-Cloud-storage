package config

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/joho/godotenv"
)

var Cfg Config = Config{}

type Config struct {
	Db    DbConfig
	MinIo MinioConfig
}

type DbConfig struct {
	User     string `config:"DB_USER"`
	Password string `config:"DB_PASSWORD"`
	Host     string `config:"DB_HOST"`
	Port     int    `config:"DB_PORT"`
	Name     string `config:"DB_NAME"`
	SSLMode  string `config:"DB_SSL"`
}

type MinioConfig struct {
	AccessKeyID     string `config:"MINIO_ACCESS_KEY_ID"`
	SecretAccessKey string `config:"MINIO_SECRET_ACCESS_KEY"`
	Endpoint        string `config:"MINIO_ENDPOINT"`
}

func LoadConfig() error {

	_, b, _, _ := runtime.Caller(0)

	projectroot := filepath.Join(filepath.Dir(b), "..")

	envpath := filepath.Join(projectroot, ".env")

	err := godotenv.Load(envpath)
	if err != nil {
		return fmt.Errorf(".env not found: %w", err)
	} else {
		log.Println(".env loaded")
	}

	loader := confita.NewLoader(env.NewBackend())

	if err := loader.Load(context.Background(), &Cfg); err != nil {
		return fmt.Errorf("load is failed: %w", err)
	}

	return nil
}
