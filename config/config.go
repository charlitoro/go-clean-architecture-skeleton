package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	// Server configuration
	ServerPort  string
	Environment string

	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// NewConfig creates a new Config instance with values from environment variables
func NewConfig() *Config {
	// Load .env file (only once, safe to call multiple times)
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found or error loading .env file")
	}

	return &Config{
		// Server configuration with defaults
		ServerPort:  getEnv("SERVER_PORT"),
		Environment: getEnv("ENVIRONMENT"),

		// Database configuration
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
	}
}

// getEnv retrieves the value of the environment variable
func getEnv(key string) string {
	return os.Getenv(key)
}

// GetDBConnectionString returns the database connection string
func (c *Config) GetDBConnectionString() string {
	// Example for MongoDB
	if c.DBUser != "" && c.DBPassword != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	}
	return fmt.Sprintf("mongodb://%s:%s/%s", c.DBHost, c.DBPort, c.DBName)
}
