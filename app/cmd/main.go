package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-on-docker/app/global"
	g "go-on-docker/app/global"
	ctl "go-on-docker/controllers"
	auth "go-on-docker/controllers/auth"
	m "go-on-docker/db/models"
	"go-on-docker/middleware"
	mw "go-on-docker/middleware"
)

func migration() error {
	if err := g.GormDB.AutoMigrate(&m.Book{}, &m.Author{}, &m.Publisher{}, &m.User{}); err != nil {
		return err
	}
	return nil
}

func success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS!",
	})
}

// 初期化関数
func init() {
	_gormDB, err := gorm.Open(mysql.Open("root:password@tcp(db:3306)/monshin?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_db, err := _gormDB.DB()
	if err != nil {
		panic(err)
	}
	g.GormDB = _gormDB
	g.Db = _db
	migration()

	global.Pem, err = ioutil.ReadFile("private.key")
	if err != nil {
		panic(err)
	}
}

var identityKey = "id"

// エントリーポイント
func main() {
	defer g.Db.Close()

	router := gin.Default()

	// Corsの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(config))

	authMiddleware, err := mw.AuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router.POST("login", authMiddleware.LoginHandler)

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/hello", auth.HelloHandle)
	}

	lobby := router.Group("/lobby")
	{
		lobby.POST("/signup", auth.Signup)
		lobby.POST("/login", auth.Login)
		lobby.GET("/validate", middleware.Authentication, auth.Validate)
	}

	// エンドポイントの設定
	router.GET("/", success)
	router.GET("/books", ctl.Book)
	router.GET("/authors", ctl.Author)
	router.GET("/author/:idx", ctl.AuthorIdx)
	router.GET("/publishers", ctl.Publisher)
	router.Run(":8000")
}
