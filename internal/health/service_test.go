package health

import (
	"fmt"
	"testing"

	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHealthRepository struct {
	mock.Mock
}

// region HealthRepository Mock

func (m *MockHealthRepository) CheckDatabase() error {
	args := m.Called()
	return args.Error(0)
}

// region HealthService Test

func TestHealthService_CheckHealth(t *testing.T) {
	tests := []struct {
		name       string
		repoError  error
		wantStatus *HealthStatus
		wantErr    bool
	}{
		{
			name:      "successful health check",
			repoError: nil,
			wantStatus: &HealthStatus{
				Database: HealthComponent{
					Status:  "healthy",
					Message: "Connected",
				},
			},
			wantErr: false,
		},
		{
			name:       "database connection error",
			repoError:  errors.NewDatabaseConnectionError(fmt.Errorf("connection failed")),
			wantStatus: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockRepo := new(MockHealthRepository)
			mockLog := new(mockLogger)

			// Configure mock behavior
			mockRepo.On("CheckDatabase").Return(tt.repoError)
			service := NewHealthService(mockRepo, mockLog)

			// Execute test
			status, err := service.CheckHealth()

			// Verify results
			if tt.wantErr {
				assert.Error(t, err)
				if tt.repoError != nil {
					assert.IsType(t, &errors.AppError{}, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantStatus, status)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
