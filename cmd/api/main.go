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
	supplierRepo := repository.NewSupplierRepository(config.DB)

	authService := service.NewAuthService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	r := gin.Default()
	apiV1Routes := r.Group("/api/v1")

	// Auth routes
	authRoutes := apiV1Routes.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)
	
	// Category Routes
	categoryRoutes := apiV1Routes.Group("/categories")
	categoryRoutes.Use(middleware.AuthMiddleware)
	categoryRoutes.GET("", categoryHandler.FindAllCategories)
	categoryRoutes.GET("/:id", categoryHandler.FindCategoryByID)
	categoryRoutes.POST("", middleware.RoleMiddleware("ADMIN"), categoryHandler.CreateCategory)
	categoryRoutes.PUT("/:id", middleware.RoleMiddleware("ADMIN"), categoryHandler.UpdateCategory)
	categoryRoutes.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), categoryHandler.DeleteCategory)

	// Supplier Routes
	supplierRoutes := apiV1Routes.Group("/suppliers")
	supplierRoutes.Use(middleware.AuthMiddleware)
	supplierRoutes.GET("", supplierHandler.FindAllSuppliers)
	supplierRoutes.GET("/:id", supplierHandler.FindSupplierByID)
	supplierRoutes.POST("", middleware.RoleMiddleware("ADMIN"), supplierHandler.CreateSupplier)
	supplierRoutes.PUT("/:id", middleware.RoleMiddleware("ADMIN"), supplierHandler.UpdateSupplier)
	supplierRoutes.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), supplierHandler.DeleteSupplier)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	r.Run(":" + port)
}