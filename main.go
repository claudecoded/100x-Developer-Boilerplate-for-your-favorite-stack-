package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"://github.com"
	"://github.com"
	"://github.com"
	"://github.com"
	"://github.com"
	"://github.com"
)

func main() {
	log.Println("Starting up the microservice orchestration bootstrap module...")
	cfg := config.LoadConfig()

	// Initialize Database and orchestrate Schema migrations setup
	db := database.NewPostgresConnection(cfg.DatabaseURL)
	defer db.Close()
	bootstrapSchema(db)

	// Clean Dependency Injection Tree Setup
	userRepo := repository.NewPostgresUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	
	router := mux.NewRouter()
	handler.NewUserHandler(router, userUC)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Channel loop setup listening to system cancellation signals for Graceful Shutdown
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening efficiently on port :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Critical: Network server failure: %v", err)
		}
	}()

	<-shutdownChannel
	log.Println("Shutdown signal received. Safely stopping microservice operations...")

	// Enforce context safety timeout for closing database connections and remaining requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
	log.Println("System components exited cleanly. Offline mode complete.")
}

func bootstrapSchema(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) UNIQUE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	);`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Fatal: Automation migration execution failed: %v", err)
	}
}
