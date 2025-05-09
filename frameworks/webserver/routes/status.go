package routes

import (
	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/gin-gonic/gin"
)

// StatusRoutes sets up the status routes
func StatusRoutes(router *gin.Engine) {
	statusController := controllers.NewStatusController()

	router.GET("/health", statusController.HealthCheck)
}
