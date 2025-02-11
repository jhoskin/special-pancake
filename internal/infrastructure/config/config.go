package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	DBPath string
	Port   string
}

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
