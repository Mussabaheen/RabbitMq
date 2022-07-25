package Queues

import (
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send(wg sync.WaitGroup) {
	defer wg.Done()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	body := "Hello World!"

	for a := 0; a <= 1000000; a++ {
		bodyToSend := body + fmt.Sprintf("%d", a)
		go func() {
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(bodyToSend),
				})
		}()

	}

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
