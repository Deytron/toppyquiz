package main

import (
	h "client/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load HTML templates folder
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	SetRoutes(r)
	r.Run(":8082")
}

func SetRoutes(r *gin.Engine) {
	r.NoRoute(h.NoRouteHandler)
	r.GET("/health", h.HealthHandler)
}
