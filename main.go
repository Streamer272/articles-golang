package main

import (
	"articles-golang/database"
	"articles-golang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Articles Golang",
	})

	app.Server().MaxConnsPerIP = 1

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
