package controllers

import (
	"articles-golang/exceptions"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		panic(fiber.ErrUnprocessableEntity)
	}

	if err := c.JSON(data); err != nil {
		panic(fiber.ErrInternalServerError)
	}

	c.Next()

	return nil
}
