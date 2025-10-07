package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bpalazzi512/easy-ballot/backend/config"
	organizationHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/organizations"
	userHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/users"
	"github.com/bpalazzi512/easy-ballot/backend/routes"
	"github.com/bpalazzi512/easy-ballot/backend/services/organizations"
	"github.com/bpalazzi512/easy-ballot/backend/services/users"
)

func main() {
	dbConfig := config.GetDatabaseConfig()

	// Connect to MongoDB
	client, db, err := config.ConnectMongoDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	userCollection := db.Collection("users")
	userRepo := users.NewMongoDBUserRepository(userCollection)

	userService := users.NewUserService(userRepo)
	userHandler := userHandler.NewHandler(userService)

	organizationCollection := db.Collection("organizations")
	organizationRepo := organizations.NewMongoDBOrganizationRepository(organizationCollection)

	organizationService := organizations.NewOrganizationService(organizationRepo)
	organizationHandler := organizationHandler.NewHandler(organizationService)

	// Setup router with middleware
	router := routes.SetupRouter()

	// Register all route groups
	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterOrganizationRoutes(router, organizationHandler)

	// Start server
	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
