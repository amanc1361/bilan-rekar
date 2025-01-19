package database

import (
	"fmt"
	"log"

	"github.com/bilan-rekar/services/user-service/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func ConnectPostgres(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db, nil
}
