package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func checkError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	count := flag.Int("count", 5000, "# count")
	flag.Parse()
	connect, err := amqp.Dial("amqp://ansible:ansible@192.168.60.10:5672/")
	checkError(err, "Failed to connect to Rabbitmq-server")
	defer connect.Close()

	ch, err := connect.Channel()
	checkError(err, "Failed to open channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my-queue", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	checkError(err, "Failed QueueDeclare")

	body := "Hello World!"
	n := 0
	for {
		err = ch.Publish(
			"",         // exchange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		checkError(err, "Failed to publish a message")
		log.Printf(" %v [x] Sent %s\n", n, body)
		n += 1
		if n > *count {
			break // exit from for
		}
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}

}
