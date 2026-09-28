package database

import (
	"database/sql"
	"log"
	"time"

	_ "://github.com"
)

// NewPostgresConnection initializes connection pooling with resilient health checks
func NewPostgresConnection(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Critical: Unable to parse Database URI: %v", err)
	}

	// Performance optimization configs for scale
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Resilient retry logic loop for container spin-ups
	for i := 1; i <= 5; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Database connection successfully established.")
			return db
		}
		log.Printf("[Retry %d/5] Database not ready yet, pausing...", i)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Critical: Database connection failed after retries: %v", err)
	return nil
}
