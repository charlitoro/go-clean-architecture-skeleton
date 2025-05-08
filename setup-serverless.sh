#!/bin/bash

# Install Serverless Framework v4.11.1
echo "Installing Serverless Framework v4.11.1..."
npm install

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "Creating sample .env file..."
    cat > .env << EOL
# AWS Lambda Environment Variables
ENVIRONMENT=development
SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=27017
DB_USER=
DB_PASSWORD=
DB_NAME=app
EOL
    echo ".env file created"
fi

echo "Setup complete!"
echo "You can now run 'make build' to build the Lambda function"
echo "Then 'make deploy' to deploy to AWS"
echo "Or 'make local' to test locally"
