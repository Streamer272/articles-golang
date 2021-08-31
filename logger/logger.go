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
	dateTime := time.Now().Format("02-01-2006 15:04:05")

	Log("Error occured: %v", dateTime, "ERROR", "0", error)
}

func LogInfo(info interface{}) {
	dateTime := time.Now().Format("02-01-2006 15:04:05")

	Log("Info: %v", dateTime, "INFO", "0", info)
}

func Log(format string, vars ...interface{}) {
	file, err := os.OpenFile("log/"+time.Now().Format("02_01_2006")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	fmt.Printf(Prefix+format+"\n", vars...)

	_, err = fmt.Fprintf(file, Prefix+format+"\n", vars...)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(Prefix+" Error (%v) occured while logging (%v)...\n", dateTime, "ERROR", "0", format, err)
		}

		if err = file.Close(); err != nil {
			fmt.Printf(Prefix+" Error (%v) occured while closing log file...\n", dateTime, "ERROR", "0", err)
		}
	}()
}
