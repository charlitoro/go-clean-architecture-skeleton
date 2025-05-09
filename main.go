package main

import (
	"log"

	"github.com/charlitoro/go-clean-architecture-skeleton/config"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/services"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize logger
	logger, err := services.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize server with logger DI
	server := webserver.NewServer(cfg.ServerPort, logger)

	logger.Info("Starting server...")
	server.Start()
}
