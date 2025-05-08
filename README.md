# GO Clean Architecture Skeleton

This project is a skeleton for a Go application following Clean Architecture principles.

## Project structure

```bash
my-go-project/
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
│       ├── gin.go  # o echo.go, fiber.go
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
├── internal/  # Para código interno no exportable
│   └── utils/
│       └── helpers.go
│
├── pkg/  # Para código que puede ser reutilizado en otros proyectos
│   └── interfaces/
│       └── repository.go
│
├── tests/
│   ├── fixtures/
│   │   └── posts.go
│   ├── integration/
│   │   └── webserver/
│   │       └── routes_test.go
│   └── unit/
│       ├── controllers/
│       │   ├── post_controller_test.go
│       │   └── status_controller_test.go
│       └── use_cases/
│           └── post_test.go
│
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── README.md
├── main.go
├── Makefile
└── .env.example
```

## Usage

### Using the Makefile

- **Build the application:**
  ```bash
  make build
  ```
- **Run the application:**
  ```bash
  make run
  ```
- **Clean build artifacts:**
  ```bash
  make clean
  ```
- **Run tests:**
  ```bash
  make test
  ```

### Manual usage

1. **Copy the example file:**
   ```bash
   cp .env.example .env
   ```
2. **Edit `.env`** with your settings.
3. **Run your app:**
   ```bash
   go run main.go
   ```
   or build and run as usual.

### Clean Architecture Aspects

- All configuration is centralized in the `config` package.
- Environment variables are loaded automatically for all parts of the app.

For more details, see the comments in `config/config.go`.
