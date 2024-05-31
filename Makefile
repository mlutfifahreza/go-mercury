build:
	go build -o ./out/ cmd/server/server.go

tidy:
	go mod tidy

test:
	go test -cover ./...

pre-commit:
	make build tidy test

check:
	make build tidy

run:
	go run cmd/server/server.go
