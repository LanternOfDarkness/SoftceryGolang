package config

type Config struct {
	BindAddr        string
	FileStoragePath string
	RabbitMQConfig  *RabbitMQConfig
}

type RabbitMQConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	QueueName string
}

func NewConfig() *Config {
	return &Config{
		BindAddr:        ":8080",
		FileStoragePath: "./images/",
		RabbitMQConfig: &RabbitMQConfig{
			Host:      "localhost",
			Port:      "5672",
			User:      "guest",
			Password:  "guest",
			QueueName: "image_queue",
		},
	}
}

func (c *Config) GetRabbitMQConfig() *RabbitMQConfig {
	return &RabbitMQConfig{
		Host:      c.RabbitMQConfig.Host,
		Port:      c.RabbitMQConfig.Port,
		User:      c.RabbitMQConfig.User,
		Password:  c.RabbitMQConfig.Password,
		QueueName: c.RabbitMQConfig.QueueName,
	}
}