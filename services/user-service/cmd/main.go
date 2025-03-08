package main

import (
	"log"

	"github.com/amanc1361/bilan-rekar/user-service/config"
	"github.com/amanc1361/bilan-rekar/user-service/internal/infrastructure/database"
	"github.com/amanc1361/bilan-rekar/user-service/internal/infrastructure/http"
)

func main() {
	// Load environment configuration

	envConfig, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Connect to the database
	db, err := database.ConnectPostgres(database.DatabaseConfig(envConfig.Database))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	defer sqlDB.Close()

	// Initialize and start the HTTP server
	http.StartServer(envConfig.AppPort, db)
}
