package main

import (
	"articles-golang/database"
	"articles-golang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Articles Golang",
	})

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
