package middleware

import (
	"articles-golang/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	fmt.Printf("path includes articles: %v\ntoken: %v\nis token valid: %v\n",
		strings.Contains(c.Path(), "articles"), c.Get("token"), services.IsTokenValid(c.Get("token")))

	if !strings.Contains(c.Path(), "articles") {
		return c.Next()
	}

	if c.Get("token") == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !services.IsTokenValid(c.Get("token")) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
