package config

import (
	"Controle-Despesas/internal/infrastructure"
	"github.com/joho/godotenv"
)

type Config struct {
	Database *database.Config
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
}
