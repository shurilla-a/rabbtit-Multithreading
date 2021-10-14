package rabbitSend

import (
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
	} else {
		defer func() {
			conn := connect.Close()
			if err == nil {
				err = conn
			} else {
				logger.ErrorLoger(err, "Не возможно закрыть connect к RebitMQ")
			}
		}()
	}
	channel, err := connect.Channel()
	if err != nil {
		logger.ErrorLoger(err, "Не возможно создать Channel для RabbitMQ")
	} else {
		//defer channel.Close()
		defer func() {
			chann := channel.Close()
			if err == nil {
				err = chann
			} else {
				logger.ErrorLoger(err, "Не возможно закрыть канал")
			}
		}()
	}

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
			if err != nil {
				logger.ErrorLoger(err, "Не возможно опубликовать сообщение")
			}

		}
	}

}
