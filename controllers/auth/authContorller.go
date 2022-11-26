package auth_contorller

import (
	"net/http"

	m "go-on-docker/db/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey string = "id"

func HelloHandle(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(http.StatusOK, gin.H{
		"user_id":   claims[identityKey],
		"user_name": user.(*m.User).UserName,
		"text":      "Hello world",
	})
}
