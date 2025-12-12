package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	// GET quiz list from sql db with GoRm
	


	// Render page with data
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}
