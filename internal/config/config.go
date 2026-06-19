// Package config loads application configuration from environment variables.
// No secrets are hard-coded here; every value is read from the environment,
// with sensible local-development defaults for non-secret settings only.
package config

import "os"

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
// These map to the existing PostgreSQL setup in the postgres/ folder.
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

// KeycloakConfig holds OIDC settings for the existing Keycloak setup
// in the keycloak/ folder.
type KeycloakConfig struct {
	IssuerURL    string
	ClientID     string
	RequiredRole string
}

// Load reads configuration from environment variables.
// See .env.example for the full list of supported variables and defaults.
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Env:  getEnv("APP_ENV", "development"),
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "app"),
			User:     getEnv("DB_USER", "app"),
			Password: getEnv("DB_PASSWORD", ""),
			// The local postgres:17-alpine substitute in postgres/compose.yml
			// runs without TLS (see DECISIONS.md "Local Development"), so
			// "disable" is the default here. Set DB_SSLMODE=require once the
			// original dhi.io/postgres image with TLS is back in use.
			SSLMode: getEnv("DB_SSLMODE", "disable"),
		},
		Keycloak: KeycloakConfig{
			IssuerURL:    getEnv("KEYCLOAK_ISSUER_URL", "http://localhost:8880/realms/workshop"),
			ClientID:     getEnv("KEYCLOAK_CLIENT_ID", "go-rest-api"),
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
