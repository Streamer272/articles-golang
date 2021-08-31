package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var articles []models.Article
	database.DB.Find(&articles)

	c.JSON(articles)

	return nil
}

func CreateArticle(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(fiber.ErrUnprocessableEntity)
	}

	var user models.User
	database.DB.Find(&user, data["userId"])

	article := models.Article{
		User:  user,
		Title: data["title"],
		Text:  data["text"],
	}

	database.DB.Create(&article)

	c.JSON(article)

	return nil
}
