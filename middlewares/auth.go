package middlewares

import (
	"articles-golang/exceptions"
	"articles-golang/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	defer exceptions.HandleException(c)

	if strings.Contains(c.Path(), "articles") && !services.IsTokenValid(c.Get("token")) {
		panic(fiber.ErrUnauthorized)
	}

	c.Next()

	return nil
}
