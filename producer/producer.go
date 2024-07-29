package producer

import (
	"scheduler/helper"

	"github.com/streadway/amqp"
)

func PublishMessage(ch *amqp.Channel, exchange string, routingKey string, body string, delay int64) {

	err := ch.ExchangeDeclare(
		exchange,                               // name
		"x-delayed-message",                    // type
		true,                                   // durable
		false,                                  // auto-deleted
		false,                                  // internal
		false,                                  // no-wait
		amqp.Table{"x-delayed-type": "direct"}, // arguments
	)
	helper.FailOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers: amqp.Table{
				"x-delay": delay, // delay in milliseconds
			},
		})
	helper.FailOnError(err, "Failed to publish a message")
}
