# Start from a Golang base image
FROM golang:1.20.4-alpine3.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build Bankie
RUN go build -o /app/bankie ./cmd/app

# Set the entrypoint to run Bankie
CMD ["/app/bankie"]
