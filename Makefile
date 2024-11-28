all: build test

build:
	@echo "Building..."
	
	@go build -o main cmd/api/main.go
run:
	@go run cmd/api/main.go
test:
	@go test ./... -v
