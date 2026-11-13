package middleware

import (
	"inventory-api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header missing or invalid",
		})
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")

	claims, err := utils.ValidateJWT(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.Set("user_id", uint(claims["user_id"].(float64)))
	c.Set("role", claims["role"])

	c.Next()
}