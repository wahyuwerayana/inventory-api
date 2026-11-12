package main

import (
	"inventory-api/config"
	"inventory-api/internal/handler"
	"inventory-api/internal/repository"
	"inventory-api/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	userRepo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	authRoutes := r.Group("/api/v1/auth")
	
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)


	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	r.Run(":" + port)
}