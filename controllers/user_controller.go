package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
	}

	database.DB.Create(&user)

	c.JSON(user)

	return nil
}

func Login(c *fiber.Ctx) error {
	// FIXME

	defer exceptions.HandleException(c)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	fmt.Printf("%v\n", user)

	if user.Id == 0 || &user == nil {
		c.Status(fiber.StatusBadRequest)
		c.SendString("User not found")
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)
		c.SendString("Incorrect password")
	}

	c.JSON(user)

	return nil
}
