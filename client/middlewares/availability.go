package middlewares

import (
	"client/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func CheckAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check backend availability via simple API request
		_, err := handlers.APICall("GET", "/health", nil)
		if err != nil {
			log.Println("Backend API is not available")
			c.AbortWithStatusJSON(503, gin.H{"error": "Service Unavailable"})
			return
		}

		c.Next()
	}
}
