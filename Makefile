.PHONY: run
run:
	swag init -g cmd/main.go
	go run cmd/main.go

.PHONY: up
up:
	docker-compose up -d
