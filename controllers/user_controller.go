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

	token, err := GenerateToken(user.Id)
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
