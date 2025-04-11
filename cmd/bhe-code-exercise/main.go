package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sarwilson417/bhe-code-exercise/go/internal/handlers"
	"github.com/sarwilson417/bhe-code-exercise/go/pkg/sieve"
)

const port = 8080

func main() {
    sieveService := sieve.NewSieve()
	h := handlers.New(sieveService)
    
    mux := http.NewServeMux()
    mux.HandleFunc("/primes", h.GetNthPrime)

    srv := &http.Server{
        Addr:    fmt.Sprintf(":%v", port),
        Handler: mux,
        ReadTimeout:    10 * time.Second, 
		WriteTimeout:   20 * time.Second, 
		IdleTimeout:    15 * time.Second,
    }

    go func() {
        log.Printf("Starting bhe-code-exercise on port %v", port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("error running server: %v\n", err)
        }
    }()

    // Create a channel to listen for termination signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    <-sigChan

    log.Println("Shutting down bhe-code-exercize...")

    // Create a context with a timeout to allow 20 seconds for requests in progress to finish
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Error shutting down gracefully: %v", err)
    }

    log.Println("Server shut down gracefully")
}