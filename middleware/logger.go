package middleware

import (
	"articles-golang/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

func LogOnMiddleWare(c *fiber.Ctx) error {
	startTime := time.Now()

	logType := "INFO"

	if c.Response().StatusCode() >= 500 {
		logType = "ERROR"
	}

	file, err := os.OpenFile("log/"+time.Now().Format("02_01_2006")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	fmt.Printf(logger.Prefix+logger.BaseMessage, dateTime, logType, time.Since(startTime), c.Route().Method, c.Path())

	_, err = fmt.Fprintf(file, logger.Prefix+logger.BaseMessage, dateTime, logType, time.Since(startTime), c.Route().Method, c.Path())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(logger.Prefix+" Error (%v) occured while logging message...\n", dateTime, "ERROR", time.Since(startTime), err)
		}

		if err = file.Close(); err != nil {
			fmt.Printf(logger.Prefix+" Error (%v) occured while closing log file...\n", dateTime, "ERROR", time.Since(startTime), err)
		}
	}()

	return nil
}
