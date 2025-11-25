package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Load HTML templates folder
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.Run(":3000")
}
