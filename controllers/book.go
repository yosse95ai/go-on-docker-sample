package contorller

import (
	g "go-on-docker/app/global"
	m "go-on-docker/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Book(c *gin.Context) {
	var b []m.Book
	if err := g.GormDB.Preload("Publisher").Preload("Authors").Find(&b).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, b)
}
