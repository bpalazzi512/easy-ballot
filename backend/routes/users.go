package routes

import (
    "github.com/gorilla/mux"
    userHandlers "github.com/bpalazzi512/easy-ballot/backend/handlers/users"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(router *mux.Router, handler *userHandlers.Handler) {
    // User CRUD endpoints
    router.HandleFunc("/users", handler.CreateUser).Methods("POST")
    router.HandleFunc("/users", handler.ListUsers).Methods("GET")
    router.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
}