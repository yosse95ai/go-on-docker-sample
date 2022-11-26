package middleware

import (
	"net/http"
	"time"

	m "go-on-docker/db/models"

	g "go-on-docker/app/global"

	auth "go-on-docker/controllers/auth"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey string = "id"

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &m.User{
		UserName: claims[identityKey].(string),
	}
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*m.User); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func authenticator(c *gin.Context) (interface{}, error) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	userName := json.UserName
	password := json.Password

	user := m.User{}
	if err := g.GormDB.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	if err := auth.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, err
	}
	return user, nil
}

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*m.User); ok {
		user := m.User{}
		if err := g.GormDB.Where("user_name = ?", v.UserName).First(&user).Error; err != nil {
			return false
		}
		c.Set("user", user)
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		},
	})
}
