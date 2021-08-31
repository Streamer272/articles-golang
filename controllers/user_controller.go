package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

const secret = "secret"

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

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		c.SendString("User not found")

		return nil
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)
		c.SendString("Incorrect password")

		return nil
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["sub"] = "1"

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	c.JSON(fiber.Map{
		"token": token,
	})

	return nil
}

func Logout(c *fiber.Ctx) error {

	return nil
}

func GetUserByToken(token string) models.User {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		panic(fiber.ErrUnauthorized)
	}

	claims := jwtToken.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return user
}
