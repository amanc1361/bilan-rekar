package http

import (
	"log"

	"github.com/bilan-rekar/services/user-service/application"
	"github.com/bilan-rekar/services/user-service/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	// Swagger

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartServer initializes and starts the HTTP server
func StartServer(port string, db *gorm.DB) {
	router := gin.Default()

	// Add Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize handlers
	userRepo := repository.NewGormUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// Handler map for dynamic route collection
	handlerMap := map[string]interface{}{
		"UserHandler": userHandler,
	}

	// Collect and register routes
	CollectRoutes(router, handlerMap)

	// Start the server
	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}
