package main

import (
	"rabbit-Multithreading/rabbitSend"
)

// функция логирования ошибок

//функция генерации строки
//func RandomString()
func main() {
	rabbitSend.RabbtiConnect("config.yml")
}
