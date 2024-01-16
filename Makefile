.PHONY: test coverage build run-dev

test:
	go test -cover -v ./...

coverage:
	go test -coverprofile=test-coverage/coverage.out ./...
	go tool cover -html=test-coverage/coverage.out -o test-coverage/coverage.html

build:
	go build -o my-app.exe ./cmd/

run-dev:
	go run ./cmd/