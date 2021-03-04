## RabbitMQ consumer

Run a consumer.

## Environment

Copy/Rename .en.example to .env<br>
Change the values of the variables to your need in .env file.<br>

## Run

Open your terminal in this current path, then run

```
go run rabbitmq-consumer.go
```

## Methods

1- `NewInstance`

- Stablish a connection with RabbitMq in pkg using the variables from .env file.<br>

2- `NewQueueConsumer`

- Opens a unique, concurrent server channel to process the bulk of AMQP messages
- Declares an exchange on the server
- Declares a queue to hold messages and deliver to consumers. Declaring creates a queue if it doesn't already exist, or ensures that an existing queue matches the same parameters.
- Binds an exchange to a queue so that publishings to the exchange will be routed to the queue when the publishing routing key matches the binding routing key.
- Immediately starts delivering queued messages.
