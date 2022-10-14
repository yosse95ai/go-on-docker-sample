package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	g "go-on-docker/app/global"
	ctl "go-on-docker/controllers"
	m "go-on-docker/db/models"
)

func migration() error {
	if err := g.GormDB.AutoMigrate(&m.Book{}, &m.Author{}, &m.Publisher{}); err != nil {
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

	e := gin.Default()
	e.GET("/", success)
	e.GET("/books", ctl.Book)
	e.GET("/authors", ctl.Author)
	e.GET("/author/:idx", ctl.AuthorIdx)
	e.GET("/publishers", ctl.Publisher)
	e.Run(":8000")
}
