package main

import (
	"github.com/deemount/gobpmn/repository"
)

func main() {

	bpmnf := repository.NewBPMNF()
	bpmnf.Set()
	err := bpmnf.Create()
	if err != nil {
		panic(err)
	}

	/*
		//AMQP & RabbitMQ
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer conn.Close()

		fmt.Println("Successfully connected to RabbitMQ Instance")

		ch, err := conn.Channel()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"TestQueue",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println(q)

		err = ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Hello World"),
			},
		)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println("Successfully published message to the queue")
	*/
}
