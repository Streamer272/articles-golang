package routes

import (
	"articles-golang/controllers"
	"articles-golang/exceptions"
	"articles-golang/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	// TODO: add check if all json fields are satisfied

	app.Use(cors.New(cors.Config{}))

	app.Get("/", func(c *fiber.Ctx) error {
		defer exceptions.HandleException(c)
		c.SendString("Welcome!")
		return nil
	})
	app.Get("/articles", controllers.GetArticles)
	app.Put("/articles", controllers.CreateArticle)

	app.Put("/users/register", controllers.Register)
	app.Post("/users/login", controllers.Login)

	app.Use(logger.LogOnMiddleWare)
}
