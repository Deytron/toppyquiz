package main

import (
	h "client/handlers"
	"client/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Check if dot env file exists and load it, crash if not present
	if _, err := os.Stat(".env"); err != nil {
		panic("Error loading .env file. Make sure it exists")
	}

	// Load HTML templates folder
	r := gin.Default()

	// Setup
	r.LoadHTMLGlob("templates/*")
	SetRoutes(r)

	// Use middlewares
	r.Use(middlewares.CheckAPI())

	r.Run(":80")
}

func SetRoutes(r *gin.Engine) {
	r.NoRoute(h.NoRouteHandler)
	r.GET("/health", h.HealthHandler)

	// Main route
	r.GET("/", h.MainHandler)
}
