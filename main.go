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
func errorLoger()
