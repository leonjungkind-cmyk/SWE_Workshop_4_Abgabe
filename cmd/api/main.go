package main

import (
	"log"

	"github.com/joho/godotenv"

	"swe-workshop-api/internal/config"
	"swe-workshop-api/internal/database"
	"swe-workshop-api/internal/server"
)

func main() {
	// .env is optional: in production, environment variables are set
	// directly and no .env file exists, so a missing file is not an error.
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := server.NewRouter(cfg, db)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
