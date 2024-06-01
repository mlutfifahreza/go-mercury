build:
	go build -o ./out/ cmd/server/server.go

lint:
	golangci-lint run

tidy:
	go mod tidy

test:
	go test -cover ./...

check:
	make build lint tidy test

run:
	go run cmd/server/server.go
