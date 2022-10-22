package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	g "go-on-docker/app/global"
	ctl "go-on-docker/controllers"
	auth "go-on-docker/controllers/auth"
	m "go-on-docker/db/models"
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
}

// エントリーポイント
func main() {
	defer g.Db.Close()

	router := gin.Default()

	// Corsの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://sample.com"}
	router.Use(cors.New(config))

	// エンドポイントの設定
	router.GET("/", success)
	router.GET("/books", ctl.Book)
	router.GET("/authors", ctl.Author)
	router.GET("/author/:idx", ctl.AuthorIdx)
	router.GET("/publishers", ctl.Publisher)
	auth_v1 := router.Group("/api/v1")
	{
		auth_v1.POST("/signup", auth.PostSignUp)
		auth_v1.POST("/signin", auth.PostSignIn)
	}
	router.Run(":8000")
}
