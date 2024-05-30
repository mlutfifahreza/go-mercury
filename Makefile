build:
	go build -o ./out/ cmd/server/server.go

tidy:
	go mod tidy

test:
	go test -cover ./...

all:
	make build tidy test