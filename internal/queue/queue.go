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

func (r *RabbitMQ) Connect() {
	config := c.NewConfig().GetRabbitMQConfig()
	r.QueueName = config.QueueName
	connAdr := "amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port
	conn, err := amqp.Dial(connAdr)
	if err != nil {
		panic(err)
	}
	r.Conn = conn
}

func (r *RabbitMQ) Close() {
	r.Conn.Close()
}

func (r *RabbitMQ) CreateChannel() {
	channel, err := r.Conn.Channel()
	if err != nil {
		panic(err)
	}
	r.Channel = channel
}

func (r *RabbitMQ) CloseChannel() {
	r.Channel.Close()
}

func (r *RabbitMQ) CreateQueue(name string) {
  queue, err := r.Channel.QueueDeclare(
    name,
    true,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    panic(err)
  }
  r.Queue = queue
}