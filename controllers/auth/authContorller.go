package auth_contorller

import (
	"errors"
	g "go-on-docker/app/global"
	m "go-on-docker/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(userId, email, name, password string) (*m.User, error) {
	if userId == "" || password == "" || email == "" || name == "" {
		err := errors.New("フォーム入力値が不正です")
		return nil, err
	}

	user := m.User{}
	g.GormDB.Where("user_id = ? OR email = ?", userId, email).First(&user)
	if user.ID != 0 {
		err := errors.New("IDもしくはEメールがすでに使用されています")
		return nil, err
	}

	encryptPw, err := PasswordEncrypt(password)
	if err != nil {
		return nil, err
	}

	user = m.User{
		UserProfile: m.UserProfile{
			UserId: userId,
			Name:   name,
			Email:  email,
		},
		Password: encryptPw,
	}
	if err := g.GormDB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func signIn(userId, password string) (*m.User, error) {
	user := m.User{}
	g.GormDB.Where("user_id = ?", userId).First(&user)
	if user.ID == 0 {
		err := errors.New("ユーザーIDが一致するアカウントが存在しません")
		return nil, err
	}
	err := CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func PostSignUp(c *gin.Context) {
	var json m.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := signUp(json.UserId, json.Email, json.Name, json.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user.Data())
}

func PostSignIn(c *gin.Context) {
	var json m.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := signIn(json.UserId, json.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user.Data())

}
