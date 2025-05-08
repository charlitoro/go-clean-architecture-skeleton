#!/bin/bash

# Check architecture
ARCH=$(uname -m)
BUILD_TARGET="build"

if [[ "$ARCH" == "x86_64" ]]; then
    echo "Detected x86_64 architecture, using x86 build target..."
    BUILD_TARGET="build-x86"
else
    echo "Using ARM build target..."
fi

# Build the Lambda function
echo "Building Lambda function..."
make $BUILD_TARGET

# Create the events directory if it doesn't exist
mkdir -p events

# Check if the event file exists
if [ ! -f events/api-gateway-event.json ]; then
    echo "Event file not found. Please create events/api-gateway-event.json"
    exit 1
fi

# Invoke the Lambda function locally
echo "Invoking Lambda function locally..."
AWS_LAMBDA_FUNCTION_NAME=test-function ./bin/main <<< "$(cat events/api-gateway-event.json)"

echo "Done!"
