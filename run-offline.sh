#!/bin/bash

# Build the Go application
echo "Building Go application..."
make build-local

# Enable debugging
export SLS_DEBUG=*

# Run serverless offline with custom parameters
echo "Starting serverless offline..."
npx serverless offline start

echo "Done!"
