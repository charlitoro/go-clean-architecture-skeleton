# GO Clean Architecture Skeleton


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
└── main.go
```

### Some suggestions

Convenciones de Go:

* Usa snake_case para nombres de archivos
* Usa CamelCase para nombres de funciones y tipos
* Usa nombres cortos y concisos

Herramientas recomendadas:

* Gin o Echo para webserver
* GORM o MongoDB driver para bases de datos
* Viper para configuración
* Zap o Logrus para logging
* Testify para testing


Aspectos de Clean Architecture:

* Mantén las dependencias dirigidas hacia adentro
* Define interfaces en capas internas
* Separa claramente dominio, casos de uso e infraestructura
