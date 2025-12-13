package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	// GET quiz list from sql db with GoRm

	// Render page with data
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}

func AdminTeamHandler(c *gin.Context) {
	// get quiz id from param
	quizId := c.Param("id")
	quizInt, err := strconv.Atoi(quizId)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quiz ID")
		return
	}

	// Get to create team, POST to add to DB
	switch c.Request.Method {
	case "POST":
		teamName := c.PostForm("team_name")
		newTeam := Team{
			Name:   teamName,
			Score:  0,
			QuizID: uint(quizInt),
		}
		Db.Create(&newTeam)
	case "GET":
		// Get quizzes for team creation form
		var quizzes []Quiz
		Db.Find(&quizzes)
	}

	// Render page with data
	c.HTML(http.StatusOK, "admin-team.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}
