# SoftceryGolang

Test task. The goal is to create an HTTP API for uploading, optimizing, and serving images.

# Packages

- `go get -u github.com/h2non/bimg` - image processing
- `go get -u github.com/gin-gonic/gin` - web framework
- `go get github.com/rabbitmq/amqp091-go` - RabbitMQ client

# Run project

## Run RabbitMQ

Run RabbitMQ with default login and password (guest/guest):

- `docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management`

or with custom login and password:

- `docker run -d --hostname my-rabbit --name some-rabbit -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management`

## Run project

- Setting up configs in `config/config.go`

- `make run` - run project

# Description

In this project, I used the following technologies:

- Golang
- RabbitMQ

# API

- `POST /upload` - upload image
- `GET /image/:id?quality=100/75/50/20` - get image by id with compression quality
