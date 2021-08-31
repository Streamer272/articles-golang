package exceptions

import (
	"articles-golang/logger"
	"github.com/gofiber/fiber/v2"
)

type ErrorJson struct {
	status int
	err    interface{}
	path   string
}

func HandleException(c *fiber.Ctx) {
	if err := recover(); err != nil {
		logger.LogError(err)

		var status int

		if err == fiber.ErrUnprocessableEntity {
			status = fiber.StatusUnprocessableEntity
		} else if err == fiber.ErrBadRequest {
			status = fiber.StatusBadRequest
		} else {
			status = fiber.StatusInternalServerError
		}

		c.SendStatus(status)
	}

	defer func() {
		if err := recover(); err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Next()
	}()
}
