package main

import (
	"inventory-api/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": "Inventory API is running",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	r.Run(":" + port)
}