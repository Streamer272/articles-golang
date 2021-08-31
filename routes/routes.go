package routes

import (
	"articles-golang/controllers"
	"articles-golang/logger"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.GetArticles)
	app.Post("/register", controllers.Register)

	app.Use(logger.LogOnMiddleWare)
}
