package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

// функция логирования ошибок
func errorLoger(errorToaddFile error, msgtoErrorLogFile string) {
	logFileWrite, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("НЕ могу создать файл лога")
	}
	log.SetOutput(logFileWrite)
	log.Printf("%s:%s", msgtoErrorLogFile, errorToaddFile)
}

//структура файла конфигурации
type ConfigYaml struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Login         string `yaml:"login"`
	Password      string `yaml:"passwd"`
	QueueName     string `yaml:"queueName"`
	QueueMessages int    `yaml:"queueMessages"`
	QueueCount    int    `yaml:"queueCount"`
	Threading     int    `yaml:"threading"`
}

// функция парсинга Yaml файла
func configYamlParsing(YamlconfigFile string) (*ConfigYaml, error) {
	YamlConfigReaderFile, err := ioutil.ReadFile(YamlconfigFile)
	if err != nil {
		errorLoger(err, "Нет или не открывается Файл")
	}
	c := &ConfigYaml{}
	err = yaml.Unmarshal(YamlConfigReaderFile, c)
	if err != nil {
		errorLoger(err, "Немогу распарсить конфиг")
	}
	return c, nil
}

//функция генерации строки
