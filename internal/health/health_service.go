package health

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
	err := s.repo.CheckDatabase()

	status := "Healthy"
	message := "Connected"
	if err != nil {
		status = "Unhealthy"
		message = err.Error()
	}

	dbstatus := &HealthStatus{
		Database: HealthComponent{
			Status:  status,
			Message: message,
		},
	}

	return dbstatus, nil
}
