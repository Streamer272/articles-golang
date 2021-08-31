package controllers

import (
	"articles-golang/database"
	"articles-golang/exceptions"
	"articles-golang/models"
	"fmt"
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

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User
	database.DB.Where("id = ?", data["userId"]).First(&user)

	article := models.Article{
		User:  user,
		Title: fmt.Sprintf("%v", data["title"]),
		Text:  fmt.Sprintf("%v", data["text"]),
	}

	database.DB.Create(&article)

	c.JSON(article)

	return nil
}
