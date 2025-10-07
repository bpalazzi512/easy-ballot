package organizations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bpalazzi512/easy-ballot/backend/services/organizations"
	"github.com/bpalazzi512/easy-ballot/backend/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	organizationService *organizations.OrganizationService
}

func NewHandler(organizationService *organizations.OrganizationService) *Handler {
	return &Handler{
		organizationService: organizationService,
	}
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var organization organizations.CreateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.organizationService.CreateOrganization(r.Context(), organization); err != nil {
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
		Message: "Organization created successfully",
		Data:    organization,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	organizationID := vars["id"]

	organization, err := h.organizationService.GetOrganizationByID(r.Context(), organizationID)
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
		Data:    organization,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetOrganizationsByOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ownerUserID := r.URL.Query().Get("owner_user_id")
	if ownerUserID == "" {
		response := types.APIResponse{
			Success: false,
			Message: "owner_user_id query parameter is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	organizations, err := h.organizationService.GetOrganizationsByOwner(r.Context(), ownerUserID)
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
		Data:    organizations,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	organizationID := vars["id"]

	var organization organizations.Organization
	if err := json.NewDecoder(r.Body).Decode(&organization); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.organizationService.UpdateOrganization(r.Context(), organizationID, organization); err != nil {
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
		Message: "Organization updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	organizationID := vars["id"]

	if err := h.organizationService.DeleteOrganization(r.Context(), organizationID); err != nil {
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
		Message: "Organization deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	limit := 10
	offset := 0

	// Parse pagination parameters
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

	organizationList, err := h.organizationService.ListOrganizations(r.Context(), limit, offset)
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
		Data:    organizationList,
	}
	json.NewEncoder(w).Encode(response)
}
