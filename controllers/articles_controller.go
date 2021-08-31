package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err := c.JSON(data); err != nil {
		return err
	}

	if err := c.Next(); err != nil {
		return err
	}

	return nil
}
