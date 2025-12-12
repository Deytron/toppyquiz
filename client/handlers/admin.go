package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	// GET

	// Render page with data
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}
