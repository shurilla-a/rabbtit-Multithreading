package logger

import (
	"log"
	"os"
)

func ErrorLoger(errorToaddFile error, msgtoErrorLogFile string) {
	logFileWrite, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("НЕ могу создать файл лога")
	}

	defer logFileWrite.Close()
	log.SetOutput(logFileWrite)
	log.Printf("%s:%s", msgtoErrorLogFile, errorToaddFile)
}
