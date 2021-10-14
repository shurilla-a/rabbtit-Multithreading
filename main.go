package main

import (
	"rabbit-Multithreading/rabbitSend"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	rabbitSend.RabbtiConnect("config.yml")
}
