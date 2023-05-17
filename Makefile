build:
	go build -o bankie cmd/app/main.go

run:
	go run cmd/app/main.go

test:
	go test -v ./...

clean:
	rm -f bankie