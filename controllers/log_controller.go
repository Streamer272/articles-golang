package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

func Log(c *fiber.Ctx) error {
	startTime := time.Now()

	logType := "INFO"

	if err := c.Next(); err != nil {
		logType = "ERROR"
	}

	file, err := os.OpenFile("log/" + time.Now().Format("02_01_2006") + ".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")
	message := "[%v %v %v] %v --> \"%v\"\n"

	_, err = fmt.Fprintf(os.Stdout, message, dateTime, logType, time.Since(startTime), c.Route().Method, c.Route().Path)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(file, message, dateTime, logType, time.Since(startTime), c.Route().Method, c.Route().Path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error (%v) occured while logging message (%v)...\n", err, message)
		}

		if err = file.Close(); err != nil {
			fmt.Printf("Error (%v) occured while closing log file...\n", err)
		}
	}()

	return nil
}
