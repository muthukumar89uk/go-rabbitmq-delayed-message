package main

import (
	"scheduler/connection"
	"scheduler/consumer"
	"scheduler/helper"
	"scheduler/router"
)

func init() {
	conn, err := connection.ConnectRabbitMQ()
	helper.FailOnError(err, "failed to connect a rabbit message queue")

	ch, err := conn.Channel()
	helper.FailOnError(err, "failed to open a channel")

	go consumer.ConsumeMessage(ch)
}

func main() {
	router.SetupRouter()
}
