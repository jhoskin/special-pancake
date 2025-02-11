package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"todo-app/internal/infrastructure/config"
	"todo-app/internal/infrastructure/db"
	"todo-app/internal/infrastructure/server"
)

func main() {
	cfg := config.New()

	// Create data directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(cfg.DBPath), 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Initialize database
	boltDB, err := db.NewBoltDB(cfg.DBPath)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer boltDB.Close()

	// Start server
	srv := server.NewServer(boltDB)
	if err := srv.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal("Server error:", err)
	}
}
