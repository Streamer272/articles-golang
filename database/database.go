package database

import (
	customLogger "articles-golang/logger"
	"articles-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	db_, err := gorm.Open(mysql.Open("articles_user:articles_password@/articles"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	db = db_

	err = db.AutoMigrate(&models.Article{}, &models.User{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			customLogger.LogError(err)
		}
	}()
}
