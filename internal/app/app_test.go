package app

// import (
// 	"testing"

// 	"github.com/null-bd/microservice-name/config"

// 	"github.com/pashagolub/pgxmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewApplication(t *testing.T) {
// 	// Create mock database connection
// 	mock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer mock.Close()

// 	// Create test config
// 	cfg := &config.Config{
// 		App: config.AppConfig{
// 			Name: "test-app",
// 			Port: 8080,
// 		},
// 		Database: config.DatabaseConfig{
// 			Host: "localhost",
// 			Port: 5432,
// 		},
// 	}

// 	// Create application
// 	app := NewApplication(mock, cfg)

// 	// Validate application
// 	assert.NotNil(t, app)
// 	assert.NotNil(t, app.Handler)
// 	assert.NotNil(t, app.DB)
// 	assert.Equal(t, cfg, app.Config)
// }
