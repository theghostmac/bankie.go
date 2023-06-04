package logger

import (
	"log"
	"os"
)

func ErrorLogs(message string) {
	ErrorLog := log.New(os.Stderr, "ERROR:\t ", log.Ltime|log.Ldate|log.Lshortfile)

	ErrorLog.Println(message)
}

func InfoLogs(message string) {
	InfoLog := log.New(os.Stdout, "INFO:\t", log.Ltime|log.Ldate)
	InfoLog.Println(message)
}
