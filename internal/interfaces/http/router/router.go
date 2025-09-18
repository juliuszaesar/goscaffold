package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliuszaesar/goscaffold/internal/interfaces/http/handler"
	"github.com/juliuszaesar/goscaffold/internal/interfaces/http/middleware"
)

// New creates a new HTTP router with all routes configured
func New(
	middlewares *middleware.Middleware,
	userHandler *handler.UserHandler,
	healthHandler *handler.HealthHandler,
) http.Handler {
	r := mux.NewRouter()

	// Apply global middleware
	r.Use(middlewares.Recovery)
	r.Use(middlewares.Logging)
	r.Use(middlewares.CORS)

	// Health check routes
	r.HandleFunc("/health", healthHandler.Health).Methods("GET")

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// User routes
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
