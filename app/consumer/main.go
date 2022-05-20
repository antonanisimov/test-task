package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func checkError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	connect, err := amqp.Dial("amqp://ansible:ansible@192.168.60.10:5672/")
	checkError(err, "Failed to connect to RabbitMQ")
	defer connect.Close()

	ch, err := connect.Channel()
	checkError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my-queue", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	checkError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	checkError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		n := 0
		for d := range msgs {
			log.Printf("%v Received a message: %s", n, d.Body)
			n = n + 1
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
