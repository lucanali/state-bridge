package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"relayer/config"
	"relayer/relayer"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create relayer instance
	r, err := relayer.NewRelayer(cfg)
	if err != nil {
		log.Fatalf("Failed to create relayer: %v", err)
	}

	// Create context that will be canceled on shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start relayer
	go func() {
		if err := r.Start(ctx); err != nil {
			log.Printf("Relayer error: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	<-sigCh
	log.Println("Shutting down relayer...")
	cancel()
}
