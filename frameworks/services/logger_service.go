package services

import (
	"os"

	"go.uber.org/zap"
)

// Logger is an interface for logging, to allow for easier testing and DI
// You can expand this interface as needed
// For now, we use zap.Logger directly

type Logger struct {
	*zap.Logger
}

// NewLogger creates a new zap.Logger instance
// Uses NewDevelopment if ENVIRONMENT=development, else NewProduction
func NewLogger() (*Logger, error) {
	env := os.Getenv("ENVIRONMENT")

	var logger *zap.Logger
	var err error

	if env == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}
	return &Logger{logger}, nil
}
