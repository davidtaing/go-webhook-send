.PHONY: help

help: 
	@echo "Available commands:"
	@echo "build: build the project"
	@echo "build-linux: build the project for linux"
	@echo "clean: remove binary"
	@echo "deps: install dependencies"
	@echo "docker-start: start docker-compose"
	@echo "docker-stop: stop docker-compose"
	@echo "lint: run linter"
	@echo "run: run the project"
	@echo "test: run tests"

send-build:
	CGO_ENABLED=0 go build -o ./build/send ./cmd/send

send-build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/send ./cmd/send

send-clean:
	rm -f ./build/send

deps:
	go mod tidy

docker-start:
	docker-compose up -d --build

docker-stop:
	docker-compose down

lint:
	golangci-lint run

run:
	go run main.go

test:
	go test ./...
