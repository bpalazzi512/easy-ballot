package routes

import (
	positionHandlers "github.com/bpalazzi512/easy-ballot/backend/handlers/positions"
	"github.com/gorilla/mux"
)

// RegisterPositionRoutes registers all position-related routes
func RegisterPositionRoutes(router *mux.Router, handler *positionHandlers.Handler) {
	// Position CRUD endpoints
	router.HandleFunc("/positions", handler.CreatePosition).Methods("POST")
	router.HandleFunc("/positions", handler.ListPositions).Methods("GET")
	router.HandleFunc("/positions/{id}", handler.GetPosition).Methods("GET")
	router.HandleFunc("/positions/{id}", handler.UpdatePosition).Methods("PUT")
	router.HandleFunc("/positions/{id}", handler.DeletePosition).Methods("DELETE")

	// Organization-specific position endpoints
	router.HandleFunc("/organizations/{organization_id}/positions", handler.GetPositionsByOrganization).Methods("GET")
}
