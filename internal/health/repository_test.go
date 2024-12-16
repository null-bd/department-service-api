package health

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
	"github.com/stretchr/testify/assert"
)

type mockLogger struct {
	logger.Logger
}

func (m *mockLogger) Debug(msg string, fields logger.Fields) {}
func (m *mockLogger) Info(msg string, fields logger.Fields)  {}
func (m *mockLogger) Error(msg string, fields logger.Fields) {}

func TestHealthRepository_CheckDatabase(t *testing.T) {
	// Setup test container
	container, connString := SetupTestContainer(t)
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	}()

	// Setup database connection
	db := SetupDatabase(t, connString)
	defer db.Close()

	mockLog := &mockLogger{}

	tests := []struct {
		name    string
		setup   func(*pgxpool.Pool)
		wantErr bool
	}{
		{
			name:    "successful database connection",
			setup:   func(db *pgxpool.Pool) {},
			wantErr: false,
		},
		{
			name: "failed database connection",
			setup: func(db *pgxpool.Pool) {
				db.Close()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create new connection for each test if previous was closed
			var testDB *pgxpool.Pool
			if tt.wantErr {
				testDB = db
			} else {
				testDB = SetupDatabase(t, connString)
				defer testDB.Close()
			}

			repo := NewHealthRepository(testDB, mockLog)
			tt.setup(testDB)

			err := repo.CheckDatabase()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
