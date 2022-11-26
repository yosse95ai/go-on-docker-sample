package auth_contorller

import (
	"go-on-docker/app/global"
	"go-on-docker/db/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Signup(c *gin.Context) {
	var request struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	password, err := PasswordEncrypt(request.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	user := models.User{
		UserName: request.UserName,
		Password: password,
	}

	if err = global.GormDB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func Login(c *gin.Context) {
	var request struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	user := models.User{}
	if err := global.GormDB.Where("user_name = ?", request.UserName).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if err := CompareHashAndPassword(user.Password, request.Password); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString(global.Pem)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.Writer.Header().Set("Jwt", tokenString)

	user.Token = &tokenString
	if err := global.GormDB.Model(&user).Update("token", tokenString).Error; err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "success"})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, user)
}
