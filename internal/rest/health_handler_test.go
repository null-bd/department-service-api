package rest_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/department-service-api/internal/health"
	"github.com/null-bd/department-service-api/internal/rest"
	"github.com/null-bd/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// region Handler Mock

type MockHealthService struct {
	mock.Mock
}

func (m *MockHealthService) CheckHealth() (*health.HealthStatus, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*health.HealthStatus), args.Error(1)
}

// MockLogger mocks the logger
type MockLogger struct {
	logger.Logger
}

func (m *MockLogger) Info(msg string, fields logger.Fields)  {}
func (m *MockLogger) Error(msg string, fields logger.Fields) {}

func setupRouter(handler rest.IHealthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handler.HealthCheck)
	return router
}

// region Handler Test

func TestHandler_HealthCheck(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockHealthService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful health check",
			setupMock: func(m *MockHealthService) {
				m.On("CheckHealth").Return(&health.HealthStatus{
					Database: health.HealthComponent{
						Status:  "healthy",
						Message: "Connected",
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "healthy",
				"details": map[string]interface{}{
					"database": map[string]interface{}{
						"status":  "healthy",
						"message": "Connected",
					},
				},
			},
		},
		{
			name: "service returns error",
			setupMock: func(m *MockHealthService) {
				m.On("CheckHealth").Return(nil, errors.New(errors.ErrDatabaseConnection, "Error", fmt.Errorf("service error")))
			},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody: map[string]interface{}{
				"code":    "MICROSERVICE_001",
				"message": "Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockSvc := new(MockHealthService)
			mockLogger := new(MockLogger)
			tt.setupMock(mockSvc)

			handler := rest.NewHandler(mockSvc, mockLogger)
			router := setupRouter(handler)

			// Create request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/health", nil)

			// Perform request
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Parse response body
			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			// Assert response body
			assert.Equal(t, tt.expectedBody, responseBody)

			// Verify that all expected mock calls were made
			mockSvc.AssertExpectations(t)
		})
	}
}
