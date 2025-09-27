package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Logger middleware for request logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "easy-ballot-backend",
	}
	
	json.NewEncoder(w).Encode(response)
}

// Example API handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := APIResponse{
		Success: true,
		Message: "API is working!",
		Data: map[string]interface{}{
			"version": "1.0.0",
			"endpoints": []string{
				"GET /health",
				"GET /api",
			},
		},
	}
	
	json.NewEncoder(w).Encode(response)
}

// Setup routes
func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	
	// Health check endpoint
	r.HandleFunc("/health", healthHandler).Methods("GET")
	
	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("", apiHandler).Methods("GET")
	
	return r
}

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Setup routes
	router := setupRoutes()
	
	// Add logging middleware
	router.Use(loggingMiddleware)
	
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Configure this for production
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	
	handler := c.Handler(router)
	
	// Start server
	log.Printf("Server starting on port %s", port)
	log.Printf("Health check available at: http://localhost:%s/health", port)
	log.Printf("API available at: http://localhost:%s/api", port)
	
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
