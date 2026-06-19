// Package config loads application configuration from environment variables.
// No secrets are hard-coded here; every value is read from the environment,
// with sensible local-development defaults for non-secret settings only.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config bundles all configuration groups needed by the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Keycloak KeycloakConfig
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Env  string
	Host string
	Port string
}

// DatabaseConfig holds PostgreSQL connection settings.
// These map to the existing PostgreSQL setup in the deployments/postgres/ folder.
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

// KeycloakConfig holds OIDC settings for the existing Keycloak setup
// in the deployments/keycloak/ folder.
type KeycloakConfig struct {
	IssuerURL    string
	ClientID     string
	RequiredRole string
}

// Load reads configuration from environment variables.
// See .env.example for the full list of supported variables and defaults.
func Load() (*Config, error) {
	// Load values from .env into the process environment.
	// If .env does not exist, the application still works with real environment variables.
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Env:  getEnv("APP_ENV", "development"),
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "kunde"),
			User:     getEnv("DB_USER", "kunde"),
			Password: getEnv("DB_PASSWORD", "p"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Keycloak: KeycloakConfig{
			IssuerURL:    getEnv("KEYCLOAK_ISSUER_URL", "http://localhost:8880/realms/javascript"),
			ClientID:     getEnv("KEYCLOAK_CLIENT_ID", "javascript-client"),
			RequiredRole: getEnv("KEYCLOAK_REQUIRED_ROLE", ""),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
