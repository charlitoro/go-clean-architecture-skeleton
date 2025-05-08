.PHONY: build run clean test

# Build the Go application
build:
	go build -o bin/main main.go

# Run the Go application (builds first if needed)
run: build
	./bin/main

# Clean build artifacts
clean:
	rm -rf bin

# Run tests
test:
	go test -v ./...
