package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

const Secret = "secret"

func Register(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	user := models.User{
		Username: fmt.Sprintf("%v", data["username"]),
		Email:    fmt.Sprintf("%v", data["email"]),
		Password: fmt.Sprintf("%v", data["password"]),
	}

	database.DB.Create(&user)

	c.JSON(user)

	return nil
}

func Login(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 || &user == nil {
		c.Status(fiber.StatusBadRequest)
		c.SendString("User not found")
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)
		c.SendString("Incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := claims.SignedString([]byte(Secret))
	if err != nil {
		panic(fiber.ErrInternalServerError)
	}

	c.JSON(fiber.Map{
		"token": token,
	})

	return nil
}

func GetUserByToken(token string) models.User {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		panic(fiber.ErrUnauthorized)
	}

	claims := jwtToken.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return user
}
