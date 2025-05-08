package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/charlitoro/go-clean-architecture-skeleton/application/use_cases/status"
	"github.com/charlitoro/go-clean-architecture-skeleton/config"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver/routes"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

// Initialize the Gin application once during cold start
func init() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Set Gin to release mode in production
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a Gin router
	router := gin.Default()

	// Initialize use cases
	healthCheckUseCase := status.NewHealthCheckUseCase()

	// Initialize controllers
	statusController := controllers.NewStatusController(healthCheckUseCase)

	// Setup routes
	routes.SetupStatusRoutes(router, statusController)

	// Initialize the Lambda adapter
	ginLambda = ginadapter.New(router)
}

// Handler function for AWS Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Check if running in Lambda environment
	if isLambda() {
		// Start the Lambda handler
		lambda.Start(Handler)
	} else {
		// Initialize configuration
		cfg := config.NewConfig()

		// Initialize use cases
		healthCheckUseCase := status.NewHealthCheckUseCase()

		// Initialize controllers
		statusController := controllers.NewStatusController(healthCheckUseCase)

		// Initialize server
		server := webserver.NewServer(cfg.ServerPort, statusController)

		// Start server
		log.Println("Starting local server...")
		server.Start()
	}
}

// isLambda determines if the code is running in AWS Lambda
func isLambda() bool {
	// Check for Lambda environment variables
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}
