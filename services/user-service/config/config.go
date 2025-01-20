package config

import (
	"fmt"
	"os"
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
	// envFile := os.Getenv("ENV_FILE")
	// if envFile == "" {
	// 	envFile = ".env" // مقدار پیش‌فرض
	// }

	// // بارگیری فایل .env
	// err := godotenv.Load(envFile)
	// if err != nil {
	// 	log.Fatalf("Error loading %s file: %v", envFile, err)
	// }
	fmt.Println(os.Getenv("DB_HOST"))
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
