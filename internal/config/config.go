package config

import (
	"Controle-Despesas/internal/infrastructure"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Database database.Config
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Database: database.Config{
			User:            os.Getenv("DB_USER"),
			Host:            os.Getenv("DB_HOST"),
			Password:        os.Getenv("DB_PASSWORD"),
			DBName:          os.Getenv("DB_NAME"),
			Port:            os.Getenv("DB_PORT"),
			ApplicationName: os.Getenv("APPLICATION_NAME"),
			SSLMode:         os.Getenv("DB_SSLMODE"),
		},
	}, nil
}
