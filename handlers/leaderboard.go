package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LeaderboardHandler(c *gin.Context) {
	// Get leaderboard id
	quizId := c.Param("id")

	// Get leaderboard from Db
	var teams []Team
	Db.Find(&teams, "quiz_id = ?", quizId)

	// Render page with data
	c.HTML(http.StatusOK, "leaderboard.tmpl", gin.H{
		"title":       "Toppyquiz - Accueil",
		"leaderboard": teams,
	})
}
