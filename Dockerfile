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
RUN go build -o bankie cmd/app/

# Use a lightweight Alpine base image
FROM alpine:latest

# Set the workdir instide the container

# Expose the port the application will listen to. Close after configuring config.InitiateConfig().
EXPOSE 8082

# Set the entrypoint to run Bankie
ENTRYPOINT ["/bankie"]