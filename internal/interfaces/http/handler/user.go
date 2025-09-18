package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/juliuszaesar/goscaffold/internal/application/dto"
	"github.com/juliuszaesar/goscaffold/internal/application/service"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusCreated, user)
}

// GetUser handles GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "User not found", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, user)
}

// UpdateUser handles PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	user, err := h.userService.UpdateUser(r.Context(), id, req)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, user)
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	if err := h.userService.DeleteUser(r.Context(), id); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Failed to delete user", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	req := dto.ListUsersRequest{
		Limit:  limit,
		Offset: offset,
	}

	users, err := h.userService.ListUsers(r.Context(), req)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to list users", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, users)
}

// writeJSONResponse writes a JSON response
func (h *UserHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse writes an error response
func (h *UserHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message, details string) {
	errorResp := dto.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Details: map[string]string{"error": details},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResp)
}
