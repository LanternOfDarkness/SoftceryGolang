package queue

import (
	"log"

	"github.com/LanternOfDarkness/SoftceryGolang/internal/image"
)


func RunReciever() error{
	rabbitMQ := NewRabbitMQ()
	rabbitMQ.Connect()
	if err := rabbitMQ.CreateChannel(); err != nil {
		return err
	}
	if err := rabbitMQ.CreateQueue(rabbitMQ.QueueName); err != nil {
		return err
	}
	var forever chan bool
	go func () error {
		if err := rabbitMQ.Receive(); err != nil {
			log.Fatal("Failed to receive message from queue:", err)
			return err
		}
		return nil
	}()
	<-forever
	
	rabbitMQ.CloseChannel()
	rabbitMQ.Close()
	return nil
}

func (r *RabbitMQ) Receive() error{
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
		return err
	}
	for msg := range msgs {

		if err:= image.ProcessImage(msg.Body); err != nil {
			return err
		}
	}
	return nil
}
