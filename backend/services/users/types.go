package users

import "time"

type UserRole string

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	OrganizationID string    `json:"organization_id"`
	ProfilePicture string    `json:"profile_picture"`
	Role           UserRole  `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserRepository interface {
	CreateUser(user User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(id string, user User) error
	DeleteUser(id string) error
	ListUsers(organizationID string, limit, offset int) ([]User, error)
	CountUsers(organizationID string) (int64, error)
}
