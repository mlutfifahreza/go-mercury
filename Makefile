build:
	go build -o ./out/ cmd/server/server.go

vet:
	go vet cmd/server/server.go

lint:
	golangci-lint run

tidy:
	go mod tidy

test:
	go test -cover ./...

pre-commit:
	make build vet lint tidy test

check:
	make build lint tidy

run:
	go run cmd/server/server.go
