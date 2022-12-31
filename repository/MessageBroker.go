package repository

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type MessageBroker struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

// NewMessageBroker
func NewMessageBroker() MessageBroker {
	return MessageBroker{}
}

//
func (broker *MessageBroker) Dial() {

	var err error

	broker.Connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected to RabbitMQ Instance")

}

//
func (broker *MessageBroker) OpenChannel() {

	var err error

	broker.Channel, err = broker.Connection.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}

}

//
func (broker *MessageBroker) DeclareQueue() {

	var err error

	broker.Queue, err = broker.Channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

}

//
func (broker *MessageBroker) Publish() {

	var err error

	err = broker.Channel.Publish(
		"",          // exchange
		"TestQueue", // routing key
		false,       // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("Hello World"),
		})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to the queue")
}
