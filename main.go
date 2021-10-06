package logger

import (
	"rabbit-Multithreading/logger"
	"rabbit-Multithreading/yamalParser"
)

// функция логирования ошибок

//функция генерации строки
//func RandomString()
func main() {

	configReader, err := yamalParser.ConfigYamlParsing("cofig.yml")
	if err != nil {
		logger.ErrorLoger(err, "Не Могу открыть файл конфигурации")
	}

}
