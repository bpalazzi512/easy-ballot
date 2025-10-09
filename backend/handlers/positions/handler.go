package positions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bpalazzi512/easy-ballot/backend/services/positions"
	"github.com/bpalazzi512/easy-ballot/backend/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	positionService *positions.PositionService
}

func NewHandler(positionService *positions.PositionService) *Handler {
	return &Handler{
		positionService: positionService,
	}
}

func (h *Handler) CreatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var position types.CreatePositionRequest
	if err := json.NewDecoder(r.Body).Decode(&position); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.positionService.CreatePosition(r.Context(), position); err != nil {
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
		Message: "Position created successfully",
		Data:    position,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	positionID := vars["id"]

	position, err := h.positionService.GetPositionByID(r.Context(), positionID)
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
		Data:    position,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	positionID := vars["id"]

	var position types.Position
	if err := json.NewDecoder(r.Body).Decode(&position); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.positionService.UpdatePosition(r.Context(), positionID, position); err != nil {
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
		Message: "Position updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeletePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	positionID := vars["id"]

	if err := h.positionService.DeletePosition(r.Context(), positionID); err != nil {
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
		Message: "Position deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ListPositions(w http.ResponseWriter, r *http.Request) {
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

	positionList, err := h.positionService.ListPositions(r.Context(), organizationID, limit, offset)
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
		Data:    positionList,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetPositionsByOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	organizationID := vars["organization_id"]

	positions, err := h.positionService.GetPositionsByOrganization(r.Context(), organizationID)
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
		Data:    positions,
	}
	json.NewEncoder(w).Encode(response)
}
