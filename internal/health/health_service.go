package health

import "github.com/null-bd/department-service-api/internal/errors"

type HealthService struct {
	repo *HealthRepository
}

func NewHealthService(repo *HealthRepository) *HealthService {
	return &HealthService{repo: repo}
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

	return &HealthStatus{
		Database: HealthComponent{
			Status:  "healthy",
			Message: "Connected",
		},
	}, nil
}
