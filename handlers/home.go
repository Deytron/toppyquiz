package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	// Render page with data
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Toppyquiz - Accueil",
	})
}
