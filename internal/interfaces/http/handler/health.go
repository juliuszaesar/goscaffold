package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juliuszaesar/goscaffold/internal/infrastructure/database"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	db *database.DB
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *database.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

// Health handles GET /health
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
		Services: map[string]string{
			"api": "ok",
		},
	}

	// Check database connection
	if err := h.db.Ping(); err != nil {
		response.Status = "degraded"
		response.Services["database"] = "error"
	} else {
		response.Services["database"] = "ok"
	}

	statusCode := http.StatusOK
	if response.Status != "ok" {
		statusCode = http.StatusServiceUnavailable
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
