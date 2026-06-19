BINARY := bin/server

.PHONY: build run test lint fmt tidy

build:
	go build -o $(BINARY) ./cmd/api

run:
	go run ./cmd/api

test:
	go test ./... -v

lint:
	golangci-lint run ./...

fmt:
	gofmt -l -w .
	goimports -l -w .

tidy:
	go mod tidy
