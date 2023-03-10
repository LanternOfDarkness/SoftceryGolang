package queue

import (
	c "github.com/LanternOfDarkness/SoftceryGolang/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
	QueueName string
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Connect() error{
	config := c.NewConfig().GetRabbitMQConfig()
	r.QueueName = config.QueueName
	connAdr := "amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port
	conn, err := amqp.Dial(connAdr)
	if err != nil {
		return err
	}
	r.Conn = conn
	return nil
}

func (r *RabbitMQ) Close() error {
	if err := r.Conn.Close(); err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) CreateChannel() error {
	channel, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	r.Channel = channel
	return nil
}

func (r *RabbitMQ) CloseChannel() error{
	if err := r.Channel.Close(); err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) CreateQueue(name string) error{
  queue, err := r.Channel.QueueDeclare(
    name,
    true,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    return err
  }
  r.Queue = queue
	return nil
}