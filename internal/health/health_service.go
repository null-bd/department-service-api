package health

import (
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

type HealthService struct {
	repo *HealthRepository
	log  logger.Logger
}

func NewHealthService(repo *HealthRepository, logger logger.Logger) *HealthService {
	return &HealthService{repo: repo, log: logger}
}

type HealthStatus struct {
	Database HealthComponent `json:"database"`
	// Add other dependencies here
}

type HealthComponent struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (s *HealthService) CheckHealth() (*HealthStatus, error) {
	s.log.Info("service : HealthCheck : begin", nil)
	if err := s.repo.CheckDatabase(); err != nil {
		// Service layer can add more context or details to the error
		if appErr, ok := err.(*errors.AppError); ok {
			appErr.WithDetails(errors.ErrorDetail{
				Field:   "database",
				Message: "Database health check failed",
			})
			return nil, appErr
		}
		return nil, err
	}

	s.log.Info("service : HealthCheck : exit", nil)
	return &HealthStatus{
		Database: HealthComponent{
			Status:  "healthy",
			Message: "Connected",
		},
	}, nil
}
