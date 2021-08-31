package controllers

import (
	"articles-golang/exceptions"
	"articles-golang/services"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	c.JSON(services.GetArticles())

	return nil
}

func GetArticle(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	c.JSON(services.GetArticle(c.Params("id")))

	return nil
}

func CreateArticle(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	c.JSON(services.CreateArticle(data["userId"], data["title"], data["text"]))

	return nil
}

func DeleteArticle(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	c.JSON(services.DeleteArticle(data["articleId"]))

	return nil
}
