.PHONY: help

help:
	@echo "Available commands:"
	@echo "deps: install dependencies"
	@echo "docker-start: start docker-compose"
	@echo "docker-stop: stop docker-compose"
	@echo "lint: run linter"
	@echo "run: run the project"
	@echo "receive-build: build receive server"
	@echo "receive-build-linux: build receive server for linux"
	@echo "receive-clean: remove receive server binaries"
	@echo "send-build: build send server"
	@echo "send-build-linux: build send server for linux"
	@echo "send-clean: remove send server binaries"
	@echo "test: run tests"

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

receive-build:
	CGO_ENABLED=0 go build -o ./build/receive ./cmd/receive

receive-build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/receive ./cmd/receive

receive-clean:
	rm -f ./build/receive

send-build:
	CGO_ENABLED=0 go build -o ./build/send ./cmd/send

send-build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/send ./cmd/send

send-clean:
	rm -f ./build/send

test:
	go test ./...
