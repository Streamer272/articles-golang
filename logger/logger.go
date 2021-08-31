package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	Prefix      = "[%v %v %v] "
	BaseMessage = "%v --> %v\n"
)

func LogError(error interface{}) {
	file, err := os.OpenFile("log/"+time.Now().Format("02_01_2006")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	fmt.Printf(Prefix+"Error occured: %v\n", dateTime, "ERROR", "", error)

	_, err = fmt.Fprintf(file, Prefix+"Error occured: %v\n", dateTime, "ERROR", "", error)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(Prefix+"Error (%v) occured while logging error (%v)...\n", dateTime, "ERROR", "0", error, err)
		}

		if err = file.Close(); err != nil {
			fmt.Printf(Prefix+"Error (%v) occured while closing log file...\n", dateTime, "ERROR", "0", err)
		}
	}()
}
