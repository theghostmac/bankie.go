# Start from a Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build Bankie
RUN go build -o bankie cmd/app/main.go

# Set the entrypoint to run Bankie
ENTRYPOINT ["./bankie"]