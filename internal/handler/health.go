// Package handler contains HTTP handlers.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Aga602/github-slack-integration/pkg/logger"
)

// HealthHandler handles health check endpoints.
type HealthHandler struct {
	log *logger.Logger
}

// NewHealthHandler creates a new health handler.
func NewHealthHandler(log *logger.Logger) *HealthHandler {
	return &HealthHandler{log: log}
}

// HealthResponse represents a health check response.
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// Health handles the /health endpoint.
func (h *HealthHandler) Health(w http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Message: "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error("failed to encode health response", "error", err)
	}
}

// Ready handles the /ready endpoint.
func (h *HealthHandler) Ready(w http.ResponseWriter, _ *http.Request) {
	// Add readiness checks here (database connection, external services, etc.)
	response := HealthResponse{
		Status:  "ok",
		Message: "Service is ready",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error("failed to encode ready response", "error", err)
	}
}
