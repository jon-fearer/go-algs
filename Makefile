build:
	go build cmd/run/main.go

run:
	go run cmd/run/main.go

test:
	go test -v ./internal
