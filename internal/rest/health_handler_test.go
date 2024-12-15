package rest

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	"github.com/null-bd/internal/health"
// )

// // Define the mock struct that embeds mock.Mock
// type MockHealthhealth struct {
// 	mock.Mock
// }

// // Implement the interface methods
// func (m *MockHealthhealth) CheckHealth() (*health.HealthStatus, error) {
// 	args := m.Called()
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).(*health.HealthStatus), args.Error(1)
// }

// func TestHealthCheck(t *testing.T) {
// 	// Set Gin to test mode
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name           string
// 		setupMock      func(*MockHealthhealth)
// 		expectedStatus int
// 		expectedBody   map[string]interface{}
// 	}{
// 		{
// 			name: "Healthy health",
// 			setupMock: func(m *MockHealthhealth) {
// 				m.On("CheckHealth").Return(&health.HealthStatus{
// 					Database: health.HealthComponent{
// 						Status:  "healthy",
// 						Message: "Connected",
// 					},
// 				}, nil)
// 			},
// 			expectedStatus: http.StatusOK,
// 			expectedBody: map[string]interface{}{
// 				"status": "healthy",
// 				"details": map[string]interface{}{
// 					"database": map[string]interface{}{
// 						"status":  "healthy",
// 						"message": "Connected",
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Unhealthy health",
// 			setupMock: func(m *MockHealthhealth) {
// 				m.On("CheckHealth").Return(nil, fmt.Errorf("database connection failed"))
// 			},
// 			expectedStatus: http.StatusOK,
// 			expectedBody: map[string]interface{}{
// 				"status": "unhealthy",
// 				"details": map[string]interface{}{
// 					"database": map[string]interface{}{
// 						"status":  "unhealthy",
// 						"message": "database connection failed",
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Setup
// 			mockSvc := new(MockHealthhealth)
// 			tt.setupMock(mockSvc)

// 			handler := NewHandler(mockSvc, nil) // Pass nil for resourcehealth since we're not testing it

// 			// Create test context
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)
// 			req := httptest.NewRequest("GET", "/health", nil)
// 			c.Request = req
// 			// Execute
// 			handler.HealthCheck(c)
// 			// Assert status code
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 			// Parse response body
// 			var response map[string]interface{}
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			// Assert response body
// 			assert.Equal(t, tt.expectedBody, response)
// 			// Verify that all expected mock calls were made
// 			mockSvc.AssertExpectations(t)
// 		})
// 	}
// }

// // TestNewHandler tests the handler constructor
// func TestNewHandler(t *testing.T) {
// 	mockHealthSvc := new(MockHealthhealth)

// 	handler := NewHandler(mockHealthSvc, nil)

// 	assert.NotNil(t, handler)
// 	assert.Equal(t, mockHealthSvc, handler.healthSvc)
// }
