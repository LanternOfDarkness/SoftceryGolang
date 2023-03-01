package queue

import (
	"github.com/LanternOfDarkness/SoftceryGolang/internal/image"
)


func RunReciever() {
	rabbitMQ := NewRabbitMQ()
	rabbitMQ.Connect()
	rabbitMQ.CreateChannel()
	rabbitMQ.CreateQueue(rabbitMQ.QueueName)
	var forever chan bool
	go func () {

		rabbitMQ.Receive()
	}()
	<-forever
	
	rabbitMQ.CloseChannel()
	rabbitMQ.Close()
}

func (r *RabbitMQ) Receive() {
	msgs, err := r.Channel.Consume(
		r.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {

		if err:= image.ProcessImage(msg.Body); err != nil {
			panic(err)
		}
	}
}
