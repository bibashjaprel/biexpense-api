package main

import (
	"log"

	"github.com/bibashjaprel/biexpense-api/internal/config"
	"github.com/bibashjaprel/biexpense-api/internal/database"
	"github.com/bibashjaprel/biexpense-api/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	defer db.Close()

	log.Println("Database connected successfully")

	r := router.Setup()

	log.Printf(
		"Starting server on http://localhost:%s",
		cfg.AppPort,
	)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
