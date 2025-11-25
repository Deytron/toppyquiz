package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
		"Title": "Page non trouvée",
	})
}

func HealthHandler(c *gin.Context) {
	c.JSON(200, nil)
}
