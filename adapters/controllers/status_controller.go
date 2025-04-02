package controllers

import (
	"net/http"

	"github.com/charlitoro/go-clean-architecture-skeleton/application/use_cases/status"
	"github.com/gin-gonic/gin"
)

// StatusController handles status-related endpoints
type StatusController struct {
	healthCheckUseCase *status.HealthCheckUseCase
}

// NewStatusController creates a new status controller
func NewStatusController(healthCheckUseCase *status.HealthCheckUseCase) *StatusController {
	return &StatusController{
		healthCheckUseCase: healthCheckUseCase,
	}
}

// HealthCheck handles the health check endpoint
func (c *StatusController) HealthCheck(ctx *gin.Context) {
	result := c.healthCheckUseCase.Execute()
	ctx.JSON(http.StatusOK, result)
}
