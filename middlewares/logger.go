package middlewares

import (
	"articles-golang/logger"
	"github.com/gofiber/fiber/v2"
	"time"
)

func LogOnMiddleWare(c *fiber.Ctx) error {
	startTime := time.Now()

	logType := "INFO"

	if c.Response().StatusCode() >= 500 {
		logType = "ERROR"
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	logger.Log(logger.BaseMessage, dateTime, logType, time.Since(startTime), c.Route().Method, c.Path())

	return c.Next()
}
