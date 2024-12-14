package database

import (
	"testing"

	"github.com/null-bd/microservice-name/config"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresConnection(t *testing.T) {
	tests := []struct {
		name    string
		config  config.DatabaseConfig
		wantErr bool
	}{
		{
			name: "Valid configuration",
			config: config.DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				User:     "test_user",
				Password: "test_pass",
				DBName:   "test_db",
				SSLMode:  "disable",
				MaxConns: 10,
				Timeout:  30,
			},
			wantErr: false,
		},
		{
			name: "Invalid port",
			config: config.DatabaseConfig{
				Host:     "localhost",
				Port:     -1,
				User:     "test_user",
				Password: "test_pass",
				DBName:   "test_db",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewPostgresConnection(&tt.config)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, db)
			assert.NotNil(t, db.Pool)

			// Test connection cleanup
			db.Close()
		})
	}
}
