package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: string(password),
	}

	database.DB.Create(&user)

	c.JSON(user)

	c.Next()
	return nil
}
