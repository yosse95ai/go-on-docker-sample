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

func UserAuthenticator(tokenString string, claims jwt.MapClaims) (interface{}, error) {
	user := models.User{}

	if err := global.GormDB.First(&user, claims["id"]).Error; err != nil {
		return models.User{}, fmt.Errorf("不正なトークンが入力されました.")
	}

	if user.Token == nil || tokenString != *user.Token {
		return models.User{}, fmt.Errorf("トークンの有効期限が切れています.")
	}
	return user, nil
}

func Authentication(authenticator func(string, jwt.MapClaims) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
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

			user, err := authenticator(tokenString, claims)
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
				return
			}

			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}
