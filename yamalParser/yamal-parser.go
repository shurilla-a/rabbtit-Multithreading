package yamalParser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"rabbit-Multithreading/logger"
)

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
	MessageLength int    `yaml:"messagesLength"`
}

// функция парсинга Yaml файла
func ConfigYamlParsing(YamlconfigFile string) (*ConfigYaml, error) {
	YamlConfigReaderFile, err := ioutil.ReadFile(YamlconfigFile)
	if err != nil {
		logger.ErrorLoger(err, "Нет или не открывается Файл")
	}
	c := &ConfigYaml{}
	err = yaml.Unmarshal(YamlConfigReaderFile, c)
	if err != nil {
		logger.ErrorLoger(err, "Немогу распарсить конфиг")
	}
	return c, nil

}
