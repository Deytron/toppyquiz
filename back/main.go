package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Load HTML templates folder
	r := gin.Default()
	r.Run(":3000")
}
