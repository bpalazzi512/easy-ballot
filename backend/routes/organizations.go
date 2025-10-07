package routes

import (
	organizationHandlers "github.com/bpalazzi512/easy-ballot/backend/handlers/organizations"
	"github.com/gorilla/mux"
)

// RegisterOrganizationRoutes registers all organization-related routes
func RegisterOrganizationRoutes(router *mux.Router, handler *organizationHandlers.Handler) {
	// Organization CRUD endpoints
	router.HandleFunc("/organizations", handler.CreateOrganization).Methods("POST")
	router.HandleFunc("/organizations", handler.ListOrganizations).Methods("GET")
	router.HandleFunc("/organizations/owner", handler.GetOrganizationsByOwner).Methods("GET")
	router.HandleFunc("/organizations/{id}", handler.GetOrganization).Methods("GET")
	router.HandleFunc("/organizations/{id}", handler.UpdateOrganization).Methods("PUT")
	router.HandleFunc("/organizations/{id}", handler.DeleteOrganization).Methods("DELETE")
}
