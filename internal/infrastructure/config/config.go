package config

import (
	"os"
	"path/filepath"
)

// Config holds the application configuration
type Config struct {
	DBPath string // Path to the database file
	Port   string // Port to listen on
}

// New creates a new Config instance with values from environment variables or defaults
func New() *Config {
	return &Config{
		DBPath: getEnvOrDefault("DB_PATH", filepath.Join("data", "todos.db")),
		Port:   getEnvOrDefault("PORT", "8080"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
