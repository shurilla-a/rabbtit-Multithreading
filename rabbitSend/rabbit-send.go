package rabbitSend

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbit-Multithreading/logger"
	"rabbit-Multithreading/randomString"
	"rabbit-Multithreading/yamalParser"
	"strconv"
)

func RabbtiConnect(fileName string) {

	configRead, err := yamalParser.ConfigYamlParsing(fileName)
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

	//		randomString.RandomString(configRead.MessageLength)
	messageCoutingQueueC := configRead.QueueMessages / configRead.QueueCount

	for i := 0; i < configRead.QueueCount; i++ {
		queueName := configRead.QueueName + strconv.Itoa(i)
		queue, err := channel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			logger.ErrorLoger(err, "НЕ могу создать очередь")
		}
		for i := 0; i < messageCoutingQueueC; i++ {
			body := randomString.RandomString(configRead.MessageLength)
			err = channel.Publish(
				"",
				queue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			fmt.Println(body)
			if err != nil {
				logger.ErrorLoger(err, "Не возможно опубликовать сообщение")
			}

		}
	}

}
