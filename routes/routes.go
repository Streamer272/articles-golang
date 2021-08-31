package routes

import (
	"articles-golang/controllers"
	"articles-golang/exceptions"
	"articles-golang/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	// TODO: add check if all json fields are satisfied

	app.Use(cors.New(cors.Config{}))
	app.Use(middlewares.CheckToken)

	app.Get("/", func(c *fiber.Ctx) error {
		defer exceptions.HandleException(c)
		c.SendString("Welcome!")
		return nil
	})

	app.Get("/articles", controllers.GetArticles)
	app.Get("/articles/:id", controllers.GetArticle)
	app.Put("/articles", controllers.CreateArticle)
	app.Delete("/articles", controllers.DeleteArticle)

	app.Put("/users/register", controllers.Register)
	app.Post("/users/login", controllers.Login)
	app.Post("/users/logout", controllers.Logout)

	app.Use(middlewares.LogOnMiddleWare)
}
