package main

import (
	h "client/handlers"
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
	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	h.InitDb()

	r.Run(":8080")
}

func SetRoutes(r *gin.Engine) {
	r.NoRoute(h.NoRouteHandler)

	// All the different routes
	r.GET("/leaderboard/:id", h.LeaderboardHandler)
	r.GET("/", h.MainHandler)

	// Admin route for managing content shown and leaderboard, with BasicAuth
	admin := r.Group("/admin")

	admin.GET("", h.AdminHandler)
	admin.GET("/add-team/:id", h.AdminTeamHandler)
	admin.POST("/add-team/:id", h.AdminTeamHandler)
	admin.GET("/add-quiz", h.AdminQuizHandler)
	admin.POST("/add-quiz", h.AdminQuizHandler)
	admin.GET("/dashboard/:id", h.AdminQuizHandler)
	admin.GET("/edit-quiz/:id", h.AdminQuizEditHandler)
	admin.POST("/edit-quiz/:id", h.AdminQuizEditHandler)
}
