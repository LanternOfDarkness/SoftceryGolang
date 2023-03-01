package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendImageToQueue(img []byte) error{
  r := NewRabbitMQ()
  if err:= r.Connect(); err != nil {
    return err
  }
  if err := r.CreateChannel(); err != nil {
    return err
  }
  if err := r.CreateQueue(r.QueueName); err != nil {
    return err
  }
  if err := r.Send(img); err != nil {
    return err
  }
  r.CloseChannel()
  r.Close()
  return nil
}

func (r *RabbitMQ) Send(img []byte) error{
  if err := r.Channel.Publish(
    "",
    r.Queue.Name,
    false,
    false,
    amqp.Publishing{
      ContentType: "text/plain",
      Body:        img,
    },
  ); err != nil {
    return err
  }
  return nil
}




