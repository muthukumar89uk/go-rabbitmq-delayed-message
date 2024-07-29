package connection

import (
	"fmt"
	"os"
	"scheduler/helper"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	err := helper.Configure(".env")
	if err != nil {
		fmt.Println("error is loading env file")
		return nil, err
	}

	host := os.Getenv("MQ_HOST")
	port := os.Getenv("MQ_PORT")
	password := os.Getenv("MQ_PASSWORD")
	user := os.Getenv("MQ_USER")

	connectionURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	Conn, err = amqp.Dial(connectionURI)
	helper.FailOnError(err, "Failed to connect to RabbitMQ")

	return Conn, nil
}
