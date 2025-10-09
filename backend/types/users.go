package types

import (
	"context"
	"time"
)

type UserRole string

type User struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	FirstName      string    `json:"first_name" bson:"first_name"`
	LastName       string    `json:"last_name" bson:"last_name"`
	Email          string    `json:"email" bson:"email"`
	Password       string    `json:"password" bson:"password"`
	OrganizationID string    `json:"organization_id" bson:"organization_id"`
	ProfilePicture string    `json:"profile_picture" bson:"profile_picture"`
	Role           UserRole  `json:"role" bson:"role"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateUserRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	OrganizationID string `json:"organization_id"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, id string, user User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, organizationID string, limit, offset int) ([]User, error)
	CountUsers(ctx context.Context, organizationID string) (int64, error)
}
