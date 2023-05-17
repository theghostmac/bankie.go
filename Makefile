build:
	go build -o bin/bankie.go

run:
	build
	./bin/bankie.go

test:
	go test -v ./...