package consumer

import (
	"log"
	"scheduler/helper"

	"github.com/streadway/amqp"
)

func ConsumeMessage(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"delayed-queue", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,             // queue name
		"delayed_key",      // routing key
		"delayed_exchange", // exchange
		false,
		nil)
	helper.FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helper.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
