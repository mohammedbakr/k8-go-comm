package main

import (
	"log"

	"github.com/k8-proxy/k8-go-comm/pkg/rabbitmq"
)

func main() {

	// Get a connection
	rabbitmqHost := "localhost"
	rabbitPort := "5672"
	rabbitUser := "guest"
	rabbitPassword := "guest"
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
