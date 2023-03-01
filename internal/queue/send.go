package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendImageToQueue(img []byte) {
  rabbitMQ := NewRabbitMQ()
  rabbitMQ.Connect()
  rabbitMQ.CreateChannel()
  rabbitMQ.CreateQueue(rabbitMQ.QueueName)
  rabbitMQ.Send(img)
  rabbitMQ.CloseChannel()
  rabbitMQ.Close()
}

func (r *RabbitMQ) Send(img []byte) {
  err := r.Channel.Publish(
    "",
    r.Queue.Name,
    false,
    false,
    amqp.Publishing{
      ContentType: "text/plain",
      Body:        img,
    },
  )
  if err != nil {
    panic(err)
  }
}




