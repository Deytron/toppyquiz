package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminQuizDashboardHandler(c *gin.Context) {

	// Render page with data
	c.HTML(http.StatusOK, "quiz-dashboard.tmpl", gin.H{
		"title": "Dashboard quiz",
	})

}
