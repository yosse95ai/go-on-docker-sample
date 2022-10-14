package main

import (
	crypto "go-on-docker/controllers/auth"
	m "go-on-docker/db/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	publisherName = "test-publisher"
	authorName1   = "test-author-1"
	authorName2   = "test-author-2"
	BookTitle1    = "test-book-1"
	BookTitle2    = "test-book-2"
)

func initAuthor(db *gorm.DB) {
	a1 := m.Author{
		Name: authorName1,
	}
	a2 := m.Author{
		Name: authorName2,
	}
	if err := db.Create(&a1).Create(&a2).Error; err != nil {
		panic(err)
	}
}

func initPublisher(db *gorm.DB) {
	p := m.Publisher{
		Name: publisherName,
	}
	if err := db.Create(&p).Error; err != nil {
		panic(err)
	}
}

func initBook(db *gorm.DB) {
	var p m.Publisher
	if err := db.First(&p).Error; err != nil {
		panic(err)
	}
	b := []m.Book{{
		Title:       BookTitle1,
		PublisherID: p.ID,
		Publisher:   p,
	}, {
		Title:       BookTitle2,
		PublisherID: p.ID,
		Publisher:   p,
	}}
	if err := db.Create(&b[0]).Create(&b[1]).Error; err != nil {
		panic(err)
	}
}

func initRelationship(db *gorm.DB) {
	var a m.Author
	if err := db.First(&a).Error; err != nil {
		panic(err)
	}
	var b m.Book
	if err := db.First(&b, 1).Error; err != nil {
		panic(err)
	}

	var b2 m.Book
	if err := db.First(&b2, 2).Error; err != nil {
		panic(err)
	}
	db.Model(&a).Association("Books").Append([]m.Book{b, b2})
	db.Save(&a)
}

func initUser(db *gorm.DB) {
	pw1, _ := crypto.PasswordEncrypt("monshin")
	pw2, _ := crypto.PasswordEncrypt("password")
	u := []m.User{{
		UserProfile: m.UserProfile{Email: "go@gmail.com", Name: "問診太郎", UserId: "123123"},
		Password:    pw1,
	}, {
		UserProfile: m.UserProfile{Email: "ibc@gmial.com", Name: "吉村弘明", UserId: "000011"},
		Password:    pw2,
	}}
	if err := db.Create(&u[0]).Create(&u[1]).Error; err != nil {
		panic(err)
	}
}

func seeds(db *gorm.DB) error {
	initAuthor(db)
	initPublisher(db)
	initBook(db)
	initRelationship(db)
	initUser(db)

	return nil
}

func main() {
	_gormDB, err := gorm.Open(mysql.Open("root:password@tcp(db:3306)/monshin?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := _gormDB.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	seeds(_gormDB)

}
