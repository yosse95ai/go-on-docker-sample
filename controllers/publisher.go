package contorller

import (
	g "go-on-docker/app/global"
	m "go-on-docker/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Publisher(c *gin.Context) {
	var b []m.Publisher
	if err := g.GormDB.Preload("Books").Find(&b).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, b)
}
