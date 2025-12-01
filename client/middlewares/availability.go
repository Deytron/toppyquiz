package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check backend availability via 

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
