package router

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"your-project/config"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewRouter(t *testing.T) {
// 	// Set Gin to test mode
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name     string
// 		config   *config.Config
// 		validate func(*testing.T, *Router)
// 	}{
// 		{
// 			name: "Default configuration",
// 			config: &config.Config{
// 				App: config.AppConfig{
// 					Env:  "development",
// 					Port: 8080,
// 				},
// 				Auth: config.AuthConfig{
// 					ServiceID: "test-service",
// 					ClientID:  "test-client",
// 				},
// 			},
// 			validate: func(t *testing.T, r *Router) {
// 				assert.NotNil(t, r.engine)
// 				assert.NotNil(t, r.authMiddleware)

// 				// Test health endpoint
// 				w := httptest.NewRecorder()
// 				req, _ := http.NewRequest("GET", "/health", nil)
// 				r.engine.ServeHTTP(w, req)

// 				assert.Equal(t, http.StatusOK, w.Code)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router, err := NewRouter(tt.config, nil)
// 			assert.NoError(t, err)
// 			assert.NotNil(t, router)

// 			if tt.validate != nil {
// 				tt.validate(t, router)
// 			}
// 		})
// 	}
// }
