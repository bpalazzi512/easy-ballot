package nominations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bpalazzi512/easy-ballot/backend/services/nominations"
	"github.com/bpalazzi512/easy-ballot/backend/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	nominationService *nominations.NominationService
}

func NewHandler(nominationService *nominations.NominationService) *Handler {
	return &Handler{
		nominationService: nominationService,
	}
}

func (h *Handler) CreateNomination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nomination types.CreateNominationRequest
	if err := json.NewDecoder(r.Body).Decode(&nomination); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.nominationService.CreateNomination(r.Context(), nomination); err != nil {
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
		Message: "Nomination created successfully",
		Data:    nomination,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetNomination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nominationID := vars["id"]

	nomination, err := h.nominationService.GetNominationByID(r.Context(), nominationID)
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
		Data:    nomination,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateNomination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nominationID := vars["id"]

	var nomination types.Nomination
	if err := json.NewDecoder(r.Body).Decode(&nomination); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.nominationService.UpdateNomination(r.Context(), nominationID, nomination); err != nil {
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
		Message: "Nomination updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateNominationStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nominationID := vars["id"]

	var statusRequest types.UpdateNominationStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&statusRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.nominationService.UpdateNominationStatus(r.Context(), nominationID, statusRequest.Status); err != nil {
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
		Message: "Nomination status updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteNomination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nominationID := vars["id"]

	if err := h.nominationService.DeleteNomination(r.Context(), nominationID); err != nil {
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
		Message: "Nomination deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ListNominations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	positionID := r.URL.Query().Get("position_id")
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

	nominationList, err := h.nominationService.ListNominations(r.Context(), positionID, limit, offset)
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
		Data:    nominationList,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetNominationsByPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	positionID := vars["position_id"]

	nominations, err := h.nominationService.GetNominationsByPosition(r.Context(), positionID)
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
		Data:    nominations,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetNominationsByNominee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nomineeID := vars["nominee_id"]

	nominations, err := h.nominationService.GetNominationsByNominee(r.Context(), nomineeID)
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
		Data:    nominations,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetNominationsByNominator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nominatorID := vars["nominator_id"]

	nominations, err := h.nominationService.GetNominationsByNominator(r.Context(), nominatorID)
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
		Data:    nominations,
	}
	json.NewEncoder(w).Encode(response)
}
