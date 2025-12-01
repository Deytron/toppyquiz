package main

import (
	h "client/handlers"
	"client/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load HTML templates folder
	r := gin.Default()

	// Setup
	r.LoadHTMLGlob("templates/*")
	SetRoutes(r)

	// Use middlewares
	r.Use(middlewares.CheckAPI)

	r.Run(":8082")
}

func SetRoutes(r *gin.Engine) {
	r.NoRoute(h.NoRouteHandler)
	r.GET("/health", h.HealthHandler)

	// Main route
	r.GET("/", h.MainHandler)
}
