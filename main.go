package main

import (
	"log"

	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/charlitoro/go-clean-architecture-skeleton/application/use_cases/status"
	"github.com/charlitoro/go-clean-architecture-skeleton/config"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize use cases
	healthCheckUseCase := status.NewHealthCheckUseCase()

	// Initialize controllers
	statusController := controllers.NewStatusController(healthCheckUseCase)

	// Initialize server
	server := webserver.NewServer(cfg.ServerPort, statusController)

	// Start server
	log.Println("Starting server...")
	server.Start()
}
