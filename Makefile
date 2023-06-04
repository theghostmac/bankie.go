.PHONY: all build run test clean

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

docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	@echo "Running Docker container..."
	@docker run --rm $(DOCKER_IMAGE_NAME)

