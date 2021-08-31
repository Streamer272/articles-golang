package database

import (
	"articles-golang/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("articles_user:articles_password@/articles"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	Db = db

	err = db.AutoMigrate(&models.Article{}, &models.User{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error occured while connecting to the database... %v\n", err)
		}
	}()
}
