package services

import (
	"articles-golang/database"
	"articles-golang/models"
	"time"
)

func CreateToken(userId uint) uint {
	token := models.Token{
		Expires: time.Now().Add(time.Hour * 2),
		UserId:  userId,
	}
	database.DB.Create(&token)

	return token.Id
}

func IsTokenValid(tokenId uint) bool {
	var token models.Token
	database.DB.Where("id = ?", tokenId).First(&token)

	return token.Expires.Unix() <= time.Now().Unix()
}

func InvalidateToken(tokenId uint) {
	var token models.Token
	database.DB.Where("id = ?", tokenId).First(&token)

	token.Expires = time.Now().Add(-time.Second)

	database.DB.Save(&token)
}
