package contorller

import (
	g "go-on-docker/app/global"
	m "go-on-docker/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Author(c *gin.Context) {
	var a []m.Author
	if err := g.GormDB.Preload("Books").Find(&a).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a)
}

func AuthorIdx(c *gin.Context) {
	id := c.Param("idx")
	var a m.Author
	if err := g.GormDB.Preload("Books").First(&a, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a)
}
