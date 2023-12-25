run:
	go mod download
	go run ./cmd/main/main.go

build:
	go build -o ./build/main ./cmd/main/main.go
