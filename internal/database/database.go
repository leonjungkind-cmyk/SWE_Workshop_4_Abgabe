// Package database sets up the GORM connection to the existing PostgreSQL
// instance (see postgres/ folder). It intentionally does not define any
// business models or run migrations yet.
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"swe-workshop-api/internal/config"
)

// Connect opens a GORM connection to PostgreSQL using the given configuration.
// sslmode always comes from cfg.SSLMode (DB_SSLMODE) — never hardcoded —
// because the existing postgres/compose.yml setup runs with TLS ("ssl=on").
func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Password, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %w", err)
	}

	return db, nil
}
