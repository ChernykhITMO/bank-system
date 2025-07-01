.PHONY: run up down build fmt swag logs clean

# Swagger генерация
swag:
	swag init -g cmd/main.go

# Запуск сервиса
run: swag
	go run cmd/main.go

# Сборка бинарника
build:
	go build -o bin/bank cmd/main.go

# Форматирование кода
fmt:
	go fmt ./...

# Поднятие docker-compose
up:
	docker-compose up -d

# Остановка контейнеров
down:
	docker-compose down

# Просмотр логов
logs:
	docker-compose logs -f

# Очистка билдов
clean:
	rm -rf bin
