package routes

import (
	nominationHandlers "github.com/bpalazzi512/easy-ballot/backend/handlers/nominations"
	"github.com/gorilla/mux"
)

// RegisterNominationRoutes registers all nomination-related routes
func RegisterNominationRoutes(router *mux.Router, handler *nominationHandlers.Handler) {
	// Nomination CRUD endpoints
	router.HandleFunc("/nominations", handler.CreateNomination).Methods("POST")
	router.HandleFunc("/nominations", handler.ListNominations).Methods("GET")
	router.HandleFunc("/nominations/{id}", handler.GetNomination).Methods("GET")
	router.HandleFunc("/nominations/{id}", handler.UpdateNomination).Methods("PUT")
	router.HandleFunc("/nominations/{id}/status", handler.UpdateNominationStatus).Methods("PATCH")
	router.HandleFunc("/nominations/{id}", handler.DeleteNomination).Methods("DELETE")

	// Position-specific nomination endpoints
	router.HandleFunc("/positions/{position_id}/nominations", handler.GetNominationsByPosition).Methods("GET")

	// User-specific nomination endpoints
	router.HandleFunc("/users/{nominee_id}/nominations", handler.GetNominationsByNominee).Methods("GET")
	router.HandleFunc("/users/{nominator_id}/nominations-made", handler.GetNominationsByNominator).Methods("GET")
}
