package health

import (
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IHealthService interface {
		CheckHealth() (*HealthStatus, error)
	}

	healthService struct {
		repo iHealthRepository
		log  logger.Logger
	}
)

// region Implementation

func NewHealthService(repo iHealthRepository, logger logger.Logger) IHealthService {
	return &healthService{repo, logger}
}

type HealthStatus struct {
	Database HealthComponent `json:"database"`
}

type HealthComponent struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (s *healthService) CheckHealth() (*HealthStatus, error) {
	s.log.Info("service : HealthCheck : begin", nil)
	if err := s.repo.CheckDatabase(); err != nil {
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
