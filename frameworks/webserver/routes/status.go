package routes

import (
	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/gin-gonic/gin"
)

// SetupStatusRoutes sets up the status routes
func SetupStatusRoutes(router *gin.Engine, statusController *controllers.StatusController) {
	router.GET("/health", statusController.HealthCheck)
}
