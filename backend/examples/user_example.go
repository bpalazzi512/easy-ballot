package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bpalazzi512/easy-ballot/backend/config"
	"github.com/bpalazzi512/easy-ballot/backend/services/users"
)

func main() {
	dbConfig := config.GetDatabaseConfig()
	client, database, err := config.ConnectMongoDB(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer config.CloseMongoDB(client)

	userCollection := database.Collection("users")
	userRepository := users.NewMongoDBUserRepository(userCollection)
	userService := users.NewUserService(userRepository)

	newUser := users.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		Password:       "securepassword123",
		OrganizationID: "org-123",
		ProfilePicture: "https://example.com/profile.jpg",
		Role:           "admin",
	}

	fmt.Println("Creating user...")
	if err := userService.CreateUser(newUser); err != nil {
		log.Printf("Failed to create user: %v", err)
	} else {
		fmt.Println("User created successfully!")
	}

	fmt.Println("\nGetting user by email...")
	user, err := userService.GetUserByEmail("john.doe@example.com")
	if err != nil {
		log.Printf("Failed to get user: %v", err)
	} else {
		fmt.Printf("Found user: %+v\n", user)
	}

	if user != nil {
		fmt.Println("\nUpdating user...")
		user.FirstName = "Jane"
		user.UpdatedAt = time.Now()

		if err := userService.UpdateUser(user.ID, *user); err != nil {
			log.Printf("Failed to update user: %v", err)
		} else {
			fmt.Println("User updated successfully!")
		}
	}

	fmt.Println("\nListing users...")
	userList, err := userService.ListUsers("org-123", 10, 0)
	if err != nil {
		log.Printf("Failed to list users: %v", err)
	} else {
		fmt.Printf("Found %d users:\n", len(userList))
		for i, u := range userList {
			fmt.Printf("  %d. %s %s (%s)\n", i+1, u.FirstName, u.LastName, u.Email)
		}
	}

	fmt.Println("\nCounting users...")
	count, err := userService.CountUsers("org-123")
	if err != nil {
		log.Printf("Failed to count users: %v", err)
	} else {
		fmt.Printf("Total users in organization: %d\n", count)
	}

	// Uncomment to test deletion
	// if user != nil {
	// 	fmt.Println("\nDeleting user...")
	// 	if err := userService.DeleteUser(user.ID); err != nil {
	// 		log.Printf("Failed to delete user: %v", err)
	// 	} else {
	// 		fmt.Println("User deleted successfully!")
	// 	}
	// }
}
