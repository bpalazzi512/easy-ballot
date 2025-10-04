# Database Operations

This document describes the MongoDB database operations implemented for the Easy Ballot application.

## Overview

The application uses MongoDB as the database with a repository pattern implementation. The database operations are organized into:

- **Config**: Database connection and configuration
- **Repository**: Data access layer with MongoDB implementation
- **Service**: Business logic layer with validation
- **Handlers**: HTTP API endpoints

## Database Configuration

The database configuration is managed through environment variables:

- `MONGODB_URI`: MongoDB connection string (default: `mongodb://localhost:27017`)
- `MONGODB_DATABASE`: Database name (default: `easy_ballot`)

## User Model

The User model includes the following fields:

```go
type User struct {
    ID             string    `json:"id"`
    FirstName      string    `json:"first_name"`
    LastName       string    `json:"last_name"`
    Email          string    `json:"email"`
    Password       string    `json:"password"`
    OrganizationID string    `json:"organization_id"`
    ProfilePicture string    `json:"profile_picture"`
    Role           string    `json:"role"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
```

## Available Operations

### Repository Interface

The `UserRepository` interface provides the following methods:

- `CreateUser(user User) error` - Create a new user
- `GetUserByID(id string) (*User, error)` - Get user by ID
- `GetUserByEmail(email string) (*User, error)` - Get user by email
- `UpdateUser(id string, user User) error` - Update existing user
- `DeleteUser(id string) error` - Delete user by ID
- `ListUsers(organizationID string, limit, offset int) ([]User, error)` - List users with pagination
- `CountUsers(organizationID string) (int64, error)` - Count users in organization

### Service Layer

The `UserService` provides business logic with validation:

- Input validation (required fields, email format, password length)
- Duplicate email checking
- Pagination parameter validation
- Error handling and meaningful error messages

## API Endpoints

The following HTTP endpoints are available:

- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user
- `GET /api/users` - List users (with query parameters: `organization_id`, `limit`, `offset`)

## Usage Examples

### Creating a User

```go
user := users.User{
    FirstName:      "John",
    LastName:       "Doe",
    Email:          "john.doe@example.com",
    Password:       "securepassword123",
    OrganizationID: "org-123",
    ProfilePicture: "https://example.com/profile.jpg",
    Role:           "admin",
}

err := userService.CreateUser(user)
```

### Getting a User

```go
// By ID
user, err := userService.GetUserByID("user-id")

// By Email
user, err := userService.GetUserByEmail("john.doe@example.com")
```

### Updating a User

```go
user.FirstName = "Jane"
user.UpdatedAt = time.Now()
err := userService.UpdateUser(user.ID, user)
```

### Listing Users

```go
// List all users in an organization
users, err := userService.ListUsers("org-123", 10, 0)

// Count users
count, err := userService.CountUsers("org-123")
```

## Error Handling

All operations return meaningful error messages:

- Validation errors for invalid input
- "User not found" for non-existent users
- "User with email already exists" for duplicate emails
- Database connection errors

## Running the Example

To run the example code:

```bash
cd backend
go run examples/user_example.go
```

Make sure MongoDB is running and accessible at the configured URI.

## Environment Setup

1. Install MongoDB locally or use MongoDB Atlas
2. Set environment variables (optional, defaults will be used):
   ```bash
   export MONGODB_URI="mongodb://localhost:27017"
   export MONGODB_DATABASE="easy_ballot"
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Dependencies

The following Go packages are used:

- `go.mongodb.org/mongo-driver` - MongoDB driver
- `github.com/gorilla/mux` - HTTP router
- `github.com/rs/cors` - CORS middleware

All dependencies are managed through `go.mod` and will be automatically downloaded when running `go mod tidy`.
