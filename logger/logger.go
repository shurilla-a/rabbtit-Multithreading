package logger

import (
	"log"
)

func ErrorLoger(errorToaddFile error, msgtoErrorLogFile string) {
	log.Printf("%s:%s", msgtoErrorLogFile, errorToaddFile)
}

func Info(msgtoErrorLogFile string) {
	log.Println(msgtoErrorLogFile)
}
