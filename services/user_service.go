package services

import (
	"articles-golang/database"
	"articles-golang/models"
	"fmt"
)

func CreateUser(username interface{}, email interface{}, password interface{}) models.User {
	user := models.User{
		Username: fmt.Sprintf("%v", username),
		Email:    fmt.Sprintf("%v", email),
		Password: fmt.Sprintf("%v", password),
	}
	database.DB.Model(&models.User{}).Create(&user)

	return user
}

func GetUser(email interface{}) models.User {
	var user models.User
	database.DB.Model(&models.User{}).Where("email = ?", email).First(&user)

	return user
}
