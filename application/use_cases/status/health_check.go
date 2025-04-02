package status

// HealthCheckUseCase represents the use case for checking the health of the service
type HealthCheckUseCase struct{}

// NewHealthCheckUseCase creates a new health check use case
func NewHealthCheckUseCase() *HealthCheckUseCase {
	return &HealthCheckUseCase{}
}

// Execute executes the health check use case
func (uc *HealthCheckUseCase) Execute() map[string]interface{} {
	return map[string]interface{}{
		"status":  "ok",
		"service": "Go Clean Architecture Service",
	}
}
