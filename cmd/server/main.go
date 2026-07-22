package main

import (
	"log"
	"os"

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

	// Pass database connection pool to router
	r := router.Setup(db)

	// Render provides the PORT environment variable.
	// Use cfg.AppPort locally as a fallback.
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.AppPort
	}

	log.Printf(
		"Starting server on port %s",
		port,
	)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
