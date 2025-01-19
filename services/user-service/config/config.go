package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type EnvConfig struct {
	AppPort  string
	Database DatabaseConfig
}

// LoadEnv reads environment variables from .env file
func LoadEnv() (*EnvConfig, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		return nil, err
	}

	return &EnvConfig{
		AppPort: os.Getenv("APP_PORT"),
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}
