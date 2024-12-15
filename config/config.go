package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LogConfig      `mapstructure:"logging"`
	Auth     AuthConfig     `mapstructure:"auth"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	MaxConns int    `mapstructure:"max_conns"`
	Timeout  int    `mapstructure:"timeout"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type AuthConfig struct {
	ServiceID    string             `mapstructure:"serviceId"`
	ClientID     string             `mapstructure:"clientId"`
	ClientSecret string             `mapstructure:"clientSecret"`
	KeycloakURL  string             `mapstructure:"keycloakUrl"`
	Realm        string             `mapstructure:"realm"`
	CacheEnabled bool               `mapstructure:"cacheEnabled"`
	CacheURL     string             `mapstructure:"cacheUrl"`
	Resources    []AuthResource     `mapstructure:"resources"`
	PublicPaths  []PublicPathConfig `mapstructure:"publicPaths"`
}

type PublicPathConfig struct {
	Path   string   `mapstructure:"path"`
	Method []string `mapstructure:"method"` // HTTP methods to bypass auth (GET, POST, etc)
}

type AuthResource struct {
	Path      string   `mapstructure:"path"`
	Method    string   `mapstructure:"method"`
	Roles     []string `mapstructure:"roles"`
	Actions   []string `mapstructure:"actions"`
	ServiceID string   `mapstructure:"serviceId,omitempty"`
}

func LoadConfig(environment string) (*Config, error) {
	var config Config

	// Initialize Viper for default config
	defaultViper := viper.New()
	defaultViper.SetConfigName("config")
	defaultViper.SetConfigType("yaml")
	defaultViper.AddConfigPath("resources")

	// Read default config
	err := defaultViper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading default config: %v", err)
	}

	// If environment is specified, load environment specific config
	if environment != "" {
		envViper := viper.New()
		envViper.SetConfigName(fmt.Sprintf("config-%s", environment))
		envViper.SetConfigType("yaml")
		envViper.AddConfigPath("resources")

		err = envViper.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("error reading environment config: %v", err)
		}

		// Merge environment config into default config
		if err := defaultViper.MergeConfigMap(envViper.AllSettings()); err != nil {
			return nil, fmt.Errorf("error merging configs: %v", err)
		}
	}

	// Set up environment variables override
	defaultViper.SetEnvPrefix("APP")
	defaultViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	defaultViper.AutomaticEnv()

	// Unmarshal the merged configuration
	if err := defaultViper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return &config, nil
}

func (a *AppConfig) GetAddress() string {
	return fmt.Sprintf(":%d", a.Port)
}

// For debugging purposes
func GetLoadedConfigFiles() []string {
	var configFiles []string

	// Check default config
	defaultPath := filepath.Join("resources", "config.yaml")
	if _, err := filepath.Abs(defaultPath); err == nil {
		configFiles = append(configFiles, defaultPath)
	}

	// Check environment specific config if APP_ENV is set
	if env := viper.GetString("APP_ENV"); env != "" {
		envPath := filepath.Join("resources", fmt.Sprintf("config-%s.yaml", env))
		if _, err := filepath.Abs(envPath); err == nil {
			configFiles = append(configFiles, envPath)
		}
	}

	return configFiles
}
