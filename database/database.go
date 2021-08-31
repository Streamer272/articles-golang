package database

import (
	customLogger "articles-golang/logger"
	"articles-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("articles_user:articles_password@/articles"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	DB = conn

	err = DB.AutoMigrate(&models.Article{}, &models.User{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			customLogger.LogError(err)
		}
	}()
}
