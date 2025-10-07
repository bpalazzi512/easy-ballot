package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bpalazzi512/easy-ballot/backend/config"
	"github.com/bpalazzi512/easy-ballot/backend/services/organizations"
)

func main() {
	dbConfig := config.GetDatabaseConfig()
	client, database, err := config.ConnectMongoDB(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer config.CloseMongoDB(client)

	organizationCollection := database.Collection("organizations")
	organizationRepository := organizations.NewMongoDBOrganizationRepository(organizationCollection)
	organizationService := organizations.NewOrganizationService(organizationRepository)
	ctx := context.Background()

	newOrganization := organizations.CreateOrganizationRequest{
		Name:        "Acme Corporation",
		Logo:        "https://example.com/logo.png",
		OwnerUserID: "user123",
	}

	fmt.Println("Creating organization...")
	if err := organizationService.CreateOrganization(ctx, newOrganization); err != nil {
		log.Printf("Failed to create organization: %v", err)
	} else {
		fmt.Println("Organization created successfully!")
	}

	fmt.Println("\nGetting organization by owner...")
	ownerOrgs, err := organizationService.GetOrganizationsByOwner(ctx, newOrganization.OwnerUserID)
	if err != nil {
		log.Printf("Failed to get organizations by owner: %v", err)
	} else {
		fmt.Printf("Found %d organizations by owner:\n", len(ownerOrgs))
		for i, org := range ownerOrgs {
			fmt.Printf("  %d. %s (ID: %s)\n", i+1, org.Name, org.ID)
		}
	}

	if len(ownerOrgs) > 0 {
		org := &ownerOrgs[0]
		fmt.Println("\nGetting organization by ID...")
		retrievedOrg, err := organizationService.GetOrganizationByID(ctx, org.ID)
		if err != nil {
			log.Printf("Failed to get organization: %v", err)
		} else {
			fmt.Printf("Retrieved organization: %+v\n", retrievedOrg)
		}

		fmt.Println("\nUpdating organization...")
		org.Name = "Acme Corporation Updated"
		org.Logo = "https://example.com/new-logo.png"
		org.UpdatedAt = time.Now()

		if err := organizationService.UpdateOrganization(ctx, org.ID, *org); err != nil {
			log.Printf("Failed to update organization: %v", err)
		} else {
			fmt.Println("Organization updated successfully!")
		}
	}

	fmt.Println("\nListing all organizations...")
	allOrgs, err := organizationService.ListOrganizations(ctx, 10, 0)
	if err != nil {
		log.Printf("Failed to list organizations: %v", err)
	} else {
		fmt.Printf("Found %d organizations:\n", len(allOrgs))
		for i, org := range allOrgs {
			fmt.Printf("  %d. %s (Owner: %s)\n", i+1, org.Name, org.OwnerUserID)
		}
	}

	fmt.Println("\nCounting organizations...")
	count, err := organizationService.CountOrganizations(ctx)
	if err != nil {
		log.Printf("Failed to count organizations: %v", err)
	} else {
		fmt.Printf("Total organizations: %d\n", count)
	}

	// Uncomment to test deletion
	if len(ownerOrgs) > 0 {
		fmt.Println("\nDeleting organization...")
		if err := organizationService.DeleteOrganization(ctx, ownerOrgs[0].ID); err != nil {
			log.Printf("Failed to delete organization: %v", err)
		} else {
			fmt.Println("Organization deleted successfully!")
		}
	}
}
