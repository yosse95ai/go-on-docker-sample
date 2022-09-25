package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "go-on-docker/db/models"
)

// DB
var gormDB *gorm.DB
var db *sql.DB

func migration() error {
	if err := gormDB.AutoMigrate(&m.Book{}, &m.Author{}, &m.Publisher{}); err != nil {
		return err
	}
	return nil
}

func success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS!!",
	})
}

func book(c *gin.Context) {
	var b []m.Book
	if err := gormDB.Preload("Publisher").Preload("Authors").Find(&b).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, b)
}

func author(c *gin.Context) {
	var a []m.Author
	if err := gormDB.Preload("Books").Find(&a).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a)
}

func authorIdx(c *gin.Context) {
	id := c.Param("idx")
	var a m.Author
	if err := gormDB.Preload("Books").First(&a, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a)
}

func publisher(c *gin.Context) {
	var b []m.Publisher
	if err := gormDB.Preload("Books").Find(&b).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, b)
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
	gormDB = _gormDB
	db = _db
	migration()
}

// エントリーポイント
func main() {
	defer db.Close()

	e := gin.Default()
	e.GET("/", success)
	e.GET("/books", book)
	e.GET("/authors", author)
	e.GET("/author/:idx", authorIdx)
	e.GET("/publishers", publisher)
	e.Run(":8000")
}
