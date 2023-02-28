package config

type Config struct {
	BindAddr         string
	RabbitMQLogin    string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
	FileStoragePath  string
}

func NewConfig() *Config {
	return &Config{
		BindAddr:         ":8080",
		RabbitMQLogin:    "guest",
		RabbitMQPassword: "guest",
		RabbitMQHost:     "localhost",
		RabbitMQPort:     "5672",
		FileStoragePath:  "./images/",
	}
}