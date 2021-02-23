package rabbitmq

import (
	"fmt"
	"net/url"

	"github.com/streadway/amqp"
)

type RMQConnection struct {
	*amqp.Connection
}

type RMQChannel struct {
	*amqp.Channel
}

func NewInstance(rabbitHost string, rabbitPort string, messagebrokeruser string, messagebrokerpassword string) (RMQConnection, error) {

	if messagebrokeruser == "" {
		messagebrokeruser = "guest"
	}

	if messagebrokerpassword == "" {
		messagebrokerpassword = "guest"
	}

	amqpUrl := url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(messagebrokeruser, messagebrokerpassword),
		Host:   fmt.Sprintf("%s:%s", rabbitHost, rabbitPort),
		Path:   "/",
	}
	conn, err := amqp.Dial(amqpUrl.String())
	if err != nil {
		return RMQConnection{}, err
	}

	return RMQConnection{conn}, err

}

func NewQueueConsumer(connection *amqp.Connection, queueName string, exchange string, routingKey string) (<-chan amqp.Delivery, error) {

	ch, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		return nil, err
	}

	consumer, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	return consumer, err

}

func NewQueuePublisher(connection *amqp.Connection, exchange string) (*amqp.Channel, error) {

	channel, err := connection.Channel()
	if err != nil {
		return RMQChannel{}, err
	}

	if err := channel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return RMQChannel{}, err
	}

	return RMQChannel{channel}, nil

}

func PublishMessage(channel *amqp.Channel, exchange string, routingKey string, message []byte) error {

	err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            message,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	)

	return err

}
