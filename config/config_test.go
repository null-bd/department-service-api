package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		envVars  map[string]string
		wantErr  bool
		validate func(*testing.T, *Config)
	}{
		{
			name: "Load default config",
			env:  "",
			validate: func(t *testing.T, cfg *Config) {
				assert.NotNil(t, cfg)
				assert.Equal(t, "development", cfg.App.Env)
			},
		},
		{
			name: "Load local config",
			env:  "local",
			validate: func(t *testing.T, cfg *Config) {
				assert.NotNil(t, cfg)
				assert.Equal(t, "local", cfg.App.Env)
			},
		},
		{
			name: "Environment variable override",
			env:  "local",
			envVars: map[string]string{
				"APP_DATABASE_PASSWORD": "test_password",
				"APP_APP_PORT":          "9090",
			},
			validate: func(t *testing.T, cfg *Config) {
				assert.Equal(t, "test_password", cfg.Database.Password)
				assert.Equal(t, 9090, cfg.App.Port)
			},
		},
		{
			name:    "Invalid config path",
			env:     "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment variables
			for k, v := range tt.envVars {
				os.Setenv(k, v)
				defer os.Unsetenv(k)
			}

			cfg, err := LoadConfig(tt.env)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.validate != nil {
				tt.validate(t, cfg)
			}
		})
	}
}
