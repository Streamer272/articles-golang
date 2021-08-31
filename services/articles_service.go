package services

import (
	"articles-golang/database"
	"articles-golang/models"
	"fmt"
)

func GetArticles() []models.Article {
	var articles []models.Article
	database.DB.Model(&models.Article{}).Find(&articles)

	return articles
}

func GetArticle(articleId interface{}) models.Article {
	var article models.Article
	database.DB.Model(&models.Article{}).Where("id = ?", articleId).First(&article)

	return article
}

func CreateArticle(userId interface{}, title interface{}, text interface{}) models.Article {
	var user models.User
	database.DB.Model(&models.User{}).Where("id = ?", userId).First(&user)

	article := models.Article{
		User:  user,
		Title: fmt.Sprintf("%v", title),
		Text:  fmt.Sprintf("%v", text),
	}

	database.DB.Create(&article)

	return article
}

func DeleteArticle(articleId interface{}) models.Article {
	var article models.Article
	database.DB.Model(&models.Article{}).Where("id = ?", articleId).First(&article)

	database.DB.Model(&models.Article{}).Where("id = ?", articleId).Delete(&models.Article{})

	return article
}
