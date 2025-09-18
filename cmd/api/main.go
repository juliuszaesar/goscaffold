package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juliuszaesar/goscaffold/internal/infrastructure/config"
	"github.com/juliuszaesar/goscaffold/internal/infrastructure/database"
	"github.com/juliuszaesar/goscaffold/internal/infrastructure/logger"
	"github.com/juliuszaesar/goscaffold/internal/infrastructure/repository"
	"github.com/juliuszaesar/goscaffold/internal/application/service"
	"github.com/juliuszaesar/goscaffold/internal/interfaces/http/handler"
	"github.com/juliuszaesar/goscaffold/internal/interfaces/http/middleware"
	"github.com/juliuszaesar/goscaffold/internal/interfaces/http/router"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New(cfg.LogLevel, cfg.Environment)
	slog.SetDefault(log)

	// Initialize database
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db, cfg.Database.MigrationsPath); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler(db)

	// Initialize middleware
	middlewares := middleware.New(log, cfg)

	// Initialize router
	r := router.New(middlewares, userHandler, healthHandler)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		slog.Info("Starting server", "port", cfg.Server.Port, "environment", cfg.Environment)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exited")
}
