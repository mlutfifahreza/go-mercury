build:
	go build -o ./out/ cmd/server/server.go

tidy:
	go mod tidy

test:
	go test -cover ./...

pre-commit:
	make build tidy test