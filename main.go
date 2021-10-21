package main

import (
	"fmt"
	"log"
	"os"
	"rabbit-Multithreading/rabbitSend"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	logFileWrite, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("НЕ могу создать файл лога")
	} else {
		defer func() {
			logFile := logFileWrite.Close()
			if err == nil {
				err = logFile
			} else {
				log.Fatalln(err, "НЕ возможно закрыть Файл")
			}
		}()
	}
	//	defer logFileWrite.Close()
	log.SetOutput(logFileWrite)

	ch := make(chan string)
	ch2 := make(chan string)
	go rabbitSend.RabbtiConnect("config.yml", ch, "1")
	go rabbitSend.RabbtiConnect("config.yml", ch2, "2")
	s := <-ch
	s2 := <-ch2
	fmt.Println(s)
	fmt.Println(s2)
	//rabbitSend.RabbtiConnect("config.yml")
}
