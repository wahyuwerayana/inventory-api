package main

import (
	"inventory-api/config"
	"inventory-api/internal/handler"
	"inventory-api/internal/middleware"
	"inventory-api/internal/repository"
	"inventory-api/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	
	userRepo := repository.NewUserRepository(config.DB)
	categoryRepo := repository.NewCategoryRepository(config.DB)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	r := gin.Default()

	// Auth routes
	authRoutes := r.Group("/api/v1/auth")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)
	
	// Category Routes
	categoryRoutes := r.Group("/api/v1/categories")
	categoryRoutes.Use(middleware.AuthMiddleware)
	categoryRoutes.GET("", categoryHandler.FindAllCategories)
	categoryRoutes.GET("/:id", categoryHandler.FindCategoryByID)
	categoryRoutes.POST("", middleware.RoleMiddleware("ADMIN"), categoryHandler.CreateCategory)
	categoryRoutes.PUT("/:id", middleware.RoleMiddleware("ADMIN"), categoryHandler.UpdateCategory)
	categoryRoutes.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), categoryHandler.DeleteCategory)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	r.Run(":" + port)
}