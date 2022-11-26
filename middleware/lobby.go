package middleware

import (
	"fmt"
	"go-on-docker/app/global"
	"go-on-docker/db/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(c *gin.Context) {
	tokenString := c.Request.Header.Get("Jwt")

	// トークンをパース
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return global.Pem, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "不正なトークンが渡されました."})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンの有効期限が切れています."})
			return
		}

		user := models.User{}

		if err := global.GormDB.First(&user, claims["id"]).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "不正なトークンが渡されました."})
			return
		}

		if tokenString != *user.Token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンの有効期限が切れています."})
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
