package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bpalazzi512/easy-ballot/backend/config"
	nominationHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/nominations"
	organizationHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/organizations"
	positionHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/positions"
	userHandler "github.com/bpalazzi512/easy-ballot/backend/handlers/users"
	"github.com/bpalazzi512/easy-ballot/backend/routes"
	"github.com/bpalazzi512/easy-ballot/backend/services/nominations"
	"github.com/bpalazzi512/easy-ballot/backend/services/organizations"
	"github.com/bpalazzi512/easy-ballot/backend/services/positions"
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

	positionCollection := db.Collection("positions")
	positionRepo := positions.NewMongoDBPositionRepository(positionCollection)

	positionService := positions.NewPositionService(positionRepo)
	positionHandler := positionHandler.NewHandler(positionService)

	nominationCollection := db.Collection("nominations")
	nominationRepo := nominations.NewMongoDBNominationRepository(nominationCollection)

	nominationService := nominations.NewNominationService(nominationRepo)
	nominationHandler := nominationHandler.NewHandler(nominationService)

	// Setup router with middleware
	router := routes.SetupRouter()

	// Register all route groups
	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterOrganizationRoutes(router, organizationHandler)
	routes.RegisterPositionRoutes(router, positionHandler)
	routes.RegisterNominationRoutes(router, nominationHandler)

	// Start server
	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
