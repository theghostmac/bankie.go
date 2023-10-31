.PHONY: all build test clean docker

BINARY_NAME := bankie
DOCKER_IMAGE_NAME := bankie

all: build

build:
	@echo "Building binary..."
	@go build -o $(BINARY_NAME) ./cmd/app

run:
	@echo "Running the application..."
	@go run ./cmd/app

test:
	go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

docker:
	@echo "Starting Docker Compose..."
	@docker-compose up

docker-down:
	@echo "Stopping Docker Compose..."
	@docker-compose down
