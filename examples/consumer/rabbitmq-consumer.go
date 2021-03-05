package main

import (
	"log"
	"os"

	"github.com/k8-proxy/k8-go-comm/pkg/rabbitmq"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	// Get a connection
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitPort := os.Getenv("RABBITMQ_PORT")
	rabbitUser := os.Getenv("RABBITMQ_USER")
	rabbitPassword := os.Getenv("RABBITMQ_PASSWORD")

	connection, err := rabbitmq.NewInstance(rabbitmqHost, rabbitPort, rabbitUser, rabbitPassword)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Start a consumer
	queueName := "icap-adaptation"
	exchange := "icap-adaptation"
	routingKey := "icap-adaptation"
	msgs, err := rabbitmq.NewQueueConsumer(connection, queueName, exchange, routingKey)

	// Consume
	for d := range msgs {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}

}
