package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Environment string
	LogLevel    string
	Server      ServerConfig
	Database    DatabaseConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	Name           string
	SSLMode        string
	MigrationsPath string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Server: ServerConfig{
			Port:         getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvAsInt("SERVER_READ_TIMEOUT", 30),
			WriteTimeout: getEnvAsInt("SERVER_WRITE_TIMEOUT", 30),
			IdleTimeout:  getEnvAsInt("SERVER_IDLE_TIMEOUT", 120),
		},
		Database: DatabaseConfig{
			Host:           getEnv("DB_HOST", "localhost"),
			Port:           getEnvAsInt("DB_PORT", 5432),
			User:           getEnv("DB_USER", "postgres"),
			Password:       getEnv("DB_PASSWORD", "postgres"),
			Name:           getEnv("DB_NAME", "goscaffold"),
			SSLMode:        getEnv("DB_SSL_MODE", "disable"),
			MigrationsPath: getEnv("DB_MIGRATIONS_PATH", "migrations"),
		},
	}

	return cfg, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvAsInt gets an environment variable as integer with a fallback value
func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
