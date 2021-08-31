package controllers

import (
	"articles-golang/database"
	"articles-golang/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secret = "secret"

func GenerateToken(userId uint) (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["sub"] = "1"

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserByToken(token string) models.User {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		panic(err)
	}

	claims := jwtToken.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return user
}

func IsAuthorized(token string) bool {
	// FIXME

	return false
}
