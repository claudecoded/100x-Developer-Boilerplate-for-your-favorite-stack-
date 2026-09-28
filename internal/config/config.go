package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

// LoadConfig fetches environment variables or injects production-ready defaults
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback local connection string for rapid development
		dbURL = "postgres://postgres:postgres123@localhost:5432/hundredx_db?sslmode=disable"
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}
}
