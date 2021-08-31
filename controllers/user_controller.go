package controllers

import (
	"articles-golang/exceptions"
	"articles-golang/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Register(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	c.JSON(services.CreateUser(data["username"], data["email"], data["password"]))

	return nil
}

func Login(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	user := services.GetUser(data["email"])

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

	token := services.CreateToken(user.Id)

	c.JSON(fiber.Map{
		"token": token,
	})

	return nil
}

func Logout(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	tokenId, _ := strconv.Atoi(fmt.Sprintf("%v", data["tokenId"]))

	services.InvalidateToken(uint(tokenId))

	c.SendStatus(fiber.StatusOK)

	return nil
}
