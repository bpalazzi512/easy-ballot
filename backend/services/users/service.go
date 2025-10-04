package users

import (
	"fmt"
	"strings"
	"time"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(user User) error {
	if err := s.validateUser(user); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingUser, err := s.repository.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return s.repository.CreateUser(user)
}

func (s *UserService) GetUserByID(id string) (*User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("user ID cannot be empty")
	}

	return s.repository.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	return s.repository.GetUserByEmail(email)
}

func (s *UserService) UpdateUser(id string, user User) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	if err := s.validateUser(user); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingUser, err := s.repository.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if user.Email != existingUser.Email {
		userWithEmail, err := s.repository.GetUserByEmail(user.Email)
		if err == nil && userWithEmail != nil {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
	}

	user.CreatedAt = existingUser.CreatedAt
	user.UpdatedAt = time.Now()

	return s.repository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	return s.repository.DeleteUser(id)
}

func (s *UserService) ListUsers(organizationID string, limit, offset int) ([]User, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.repository.ListUsers(organizationID, limit, offset)
}

func (s *UserService) CountUsers(organizationID string) (int64, error) {
	return s.repository.CountUsers(organizationID)
}

func (s *UserService) validateUser(user User) error {
	if strings.TrimSpace(user.FirstName) == "" {
		return fmt.Errorf("first name is required")
	}
	if strings.TrimSpace(user.LastName) == "" {
		return fmt.Errorf("last name is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return fmt.Errorf("email is required")
	}
	if !isValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}
	if strings.TrimSpace(user.Password) == "" {
		return fmt.Errorf("password is required")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	if strings.TrimSpace(user.OrganizationID) == "" {
		return fmt.Errorf("organization ID is required")
	}
	if strings.TrimSpace(string(user.Role)) == "" {
		return fmt.Errorf("role is required")
	}

	return nil
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
