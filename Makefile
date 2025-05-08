.PHONY: clean deploy dev-deploy prod-deploy local format test build-local

# Clean temporary files and build artifacts
clean:
	rm -rf ./bin ./.serverless

# Build for local development with serverless-offline
build-local:
	mkdir -p bin
	go build -o bin/main main.go

# Deploy using serverless framework (serverless-go-plugin will handle the build)
deploy:
	serverless deploy --verbose

# Deploy to specific environments
dev-deploy:
	serverless deploy --stage dev --verbose

prod-deploy:
	serverless deploy --stage prod --verbose

# Run serverless offline for local development
local: build-local
	SLS_DEBUG=* serverless offline start

# Test the application
test:
	go test -v ./...

# Format the code
format:
	go fmt ./...
