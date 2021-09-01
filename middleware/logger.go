package middleware

import (
	"articles-golang/logger"
	"github.com/gofiber/fiber/v2"
	"time"
)

func LogOnMiddleWare(c *fiber.Ctx) error {
	startTime := time.Now()

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	logger.Log(logger.BaseMessage, dateTime, "INFO", time.Since(startTime), c.Route().Method, c.Path())

	return nil
}
