package middleware

import (
	"articles-golang/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	if strings.Contains(c.Path(), "articles") && !services.IsTokenValid(c.Get("token")) {
		c.SendStatus(fiber.StatusUnauthorized)

		return nil
	}

	return c.Next()
}
