package rabbitSend

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbit-Multithreading/logger"
	"rabbit-Multithreading/randomString"
	"rabbit-Multithreading/yamalParser"
)

func RabbtiConnect(multithreading int) {

	configRead, err := yamalParser.ConfigYamlParsing("config.yml")
	if err != nil {
		logger.ErrorLoger(err, "НЕ - корректный файл Конфигурации")
	}
	connectUrl := "amqp://" + configRead.Login + ":" + configRead.Password + "@" + configRead.Host + ":" + configRead.Port
	connect, err := amqp.Dial(connectUrl)
	if err != nil {
		logger.ErrorLoger(err, "НЕ возможно подключиться к RabbitMQ")
	}
	defer connect.Close()

	channel, err := connect.Channel()
	if err != nil {
		logger.ErrorLoger(err, "Не возможно создать Channel для RabbitMQ")
	}
	defer channel.Close()

	randomString.RandomString(configRead.MessageLength)

}
