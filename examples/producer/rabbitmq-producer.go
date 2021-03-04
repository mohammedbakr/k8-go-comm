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

	// Start a producer
	exchange := "icap-adaptation"
	routingKey := "icap-adaptation"
	publisher, err := rabbitmq.NewQueuePublisher(connection, exchange)

	// Publish a message
	err = rabbitmq.PublishMessage(publisher, exchange, routingKey, nil, []byte("test"))
	if err != nil {
		log.Fatalf("%s", err)
	}

}
