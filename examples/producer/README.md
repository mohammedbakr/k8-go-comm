## RabbitMQ producer

Run a producer and publish a message contains the word "test".

## Environment

Copy/Rename .en.example to .env<br>
Change the values of the variables to your need in .env file.<br>

## Run

Open your terminal in this current path, then run

```
go run rabbitmq-producer.go
```

## Methods

1- `NewInstance`

- Stablish a connection with RabbitMq in pkg using the variables from .env file.<br>

2- `NewQueuePublisher`

- Opens a unique, concurrent server channel to process the bulk of AMQP messages.
- Declares an exchange on the server.<br>

3- `PublishMessage`

- Sends a Publishing from the client to an exchange on the server.
