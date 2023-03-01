run:
	go run ./cmd/image-api/main.go

docker:
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management