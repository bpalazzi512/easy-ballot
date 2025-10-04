package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bpalazzi512/easy-ballot/backend/types"
	"github.com/bpalazzi512/easy-ballot/backend/services/users"
	"github.com/gorilla/mux"
)

type Handler struct {
	userService *users.UserService
}

func NewHandler(userService *users.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.userService.CreateUser(user); err != nil {
		response := types.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := types.APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		response := types.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := types.APIResponse{
		Success: true,
		Data:    user,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["id"]

	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.userService.UpdateUser(userID, user); err != nil {
		response := types.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := types.APIResponse{
		Success: true,
		Message: "User updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["id"]

	if err := h.userService.DeleteUser(userID); err != nil {
		response := types.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := types.APIResponse{
		Success: true,
		Message: "User deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	organizationID := r.URL.Query().Get("organization_id")
	limit := 10
	offset := 0

	// Parse pagination parameters (basic implementation)
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	if o := r.URL.Query().Get("offset"); o != "" {
		if parsedOffset, err := strconv.Atoi(o); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	userList, err := h.userService.ListUsers(organizationID, limit, offset)
	if err != nil {
		response := types.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := types.APIResponse{
		Success: true,
		Data:    userList,
	}
	json.NewEncoder(w).Encode(response)
}
