package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	shell "fluxo-go/lancamentos/shell"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize router with middleware
	router := chi.NewRouter()

	// Add standard middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(15 * time.Second))

	// Initialize Huma API
	api := humachi.New(
		router,
		huma.DefaultConfig(
			"Fluxo API",
			"1.0.0",
		),
	)

	// Register POST /lancamentos endpoint
	huma.Post(
		api,
		"/lancamentos",
		handleEfetuarLancamento,
	)

	// Health check endpoint
	huma.Get(
		api,
		"/health",
		handleHealth,
	)

	// Server configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Starting server on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server shutdown error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped")
}

// handleEfetuarLancamento handles POST /lancamentos requests
func handleEfetuarLancamento(
	ctx context.Context,
	input *shell.EfetuarLancamentoRequest,
) (
	*shell.LancamentoEfetuadoResponse,
	error,
) {
	result := input.Handle()

	if result.IsError() {
		return nil, result.UnwrapError()
	}

	response := result.Unwrap()
	return &response, nil
}

// healthResponse represents the health check response
type healthResponse struct {
	Status string `json:"status" example:"healthy"`
}

// handleHealth handles GET /health requests
func handleHealth(ctx context.Context) (*healthResponse, error) {
	return &healthResponse{Status: "healthy"}, nil
}
