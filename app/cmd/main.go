package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "SUCCESS",
		})
	})
	e.Run(":8000")
}
