start: up build

build:
	go build -v ./cmd/apiserver && \
	./apiserver

test:
	go test -v -race -timeout 30s ./...

up:
	docker-compose -f ./deployment/docker-compose.yml up -d

down:
	docker-compose -f ./deployment/docker-compose.yml down

migrate-up:
	migrate -path ./migrations -database "postgres://task:task@localhost:5432/restapi_dev?sslmode=disable" up

.DEFAULT_GOAL := build