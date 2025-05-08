# GO Clean Architecture Skeleton with AWS Lambda

This project is a skeleton for a Go application following the Clean Architecture principles, designed to be deployed as an AWS Lambda function using the Serverless Framework v4.11.1.

## Project structure

```bash
go-clean-architecture-skeleton/
│
├── adapters/
│   └── controllers/
│       ├── auth_controller.go
│       ├── post_controller.go
│       ├── status_controller.go
│       └── user_controller.go
│
├── application/
│   ├── repositories/
│   │   ├── post_db_repository.go
│   │   ├── post_redis_repository.go
│   │   └── user_db_repository.go
│   ├── services/
│   │   └── auth_service.go
│   └── use_cases/
│       ├── auth/
│       │   └── login.go
│       ├── post/
│       │   ├── add.go
│       │   ├── count_all.go
│       │   ├── delete_by_id.go
│       │   ├── find_all.go
│       │   ├── find_by_id.go
│       │   └── update_by_id.go
│       ├── status/
│       │   └── health_check.go
│       └── user/
│           ├── add.go
│           ├── count_all.go
│           ├── find_by_id.go
│           └── find_by_property.go
│
├── config/
│   └── config.go
│
├── domain/
│   └── entities/
│       ├── post.go
│       └── user.go
│
├── frameworks/
│   ├── database/
│   │   ├── mongodb/
│   │   │   ├── connection.go
│   │   │   ├── models/
│   │   │   │   ├── post.go
│   │   │   │   └── user.go
│   │   │   └── repositories/
│   │   │       ├── post_repository_mongodb.go
│   │   │       └── user_repository_mongodb.go
│   │   └── redis/
│   │       ├── connection.go
│   │       └── post_repository_redis.go
│   ├── services/
│   │   ├── auth_service.go
│   │   └── logger_service.go
│   └── webserver/
│       ├── gin.go  # or echo.go, fiber.go
│       ├── middlewares/
│       │   ├── auth_middleware.go
│       │   ├── error_handling_middleware.go
│       │   └── redis_caching_middleware.go
│       ├── routes/
│       │   ├── auth.go
│       │   ├── index.go
│       │   ├── post.go
│       │   ├── status.go
│       │   └── user.go
│       └── server.go
│
├── events/
│   └── api-gateway-event.json
│
├── bin/
│   └── main
│
├── Makefile
├── package.json
├── serverless.yml
├── setup-serverless.sh
├── test-local.sh
├── go.mod
├── go.sum
├── .env
├── .gitignore
└── README.md
```

## Prerequisites

- Go 1.x
- Node.js v16+ and npm (for Serverless Framework)
- AWS CLI configured with appropriate credentials

## Setup

1. Clone this repository

2. Install the Serverless Framework and dependencies:

```bash
./setup-serverless.sh
```

## Local Development

To run the application locally:

```bash
go run main.go
```

To test the Lambda function locally:

```bash
./test-local.sh
```

To use the Serverless Offline plugin for local testing:

```bash
make local
```

Or using npm:

```bash
npm run local
```

## Building and Deploying

1. Build the Lambda function:

```bash
make build
```

2. Deploy to AWS:

```bash
make deploy    # Default stage
make dev-deploy  # Explicitly deploy to dev stage
make prod-deploy # Deploy to production stage
```

Or using npm:

```bash
npm run deploy  # Default
npm run dev     # Dev stage
npm run prod    # Production stage
```

## Go Conventions

- Use snake_case for file names
- Use CamelCase for function and type names
- Use short and concise names

## Recommended Tools

- Gin for web server
- GORM or MongoDB driver for databases
- Viper for configuration
- Zap or Logrus for logging
- Testify for testing

## Clean Architecture Aspects

- Keep dependencies pointing inward
- Define interfaces in inner layers
- Clearly separate domain, use cases, and infrastructure

## AWS Lambda Integration

The project is configured to run as an AWS Lambda function with API Gateway using Serverless Framework v4.11.1:

- Uses AWS Lambda on ARM64 architecture (cost-efficient)
- Runs on Amazon Linux 2023 (provided.al2023 runtime)
- Supports both local development and AWS deployment
- Configured for HTTP API Gateway
- Includes Serverless Offline for local testing

The main.go file handles both local server mode and Lambda mode, automatically detecting the environment.
