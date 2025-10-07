package organizations

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type OrganizationService struct {
	repository OrganizationRepository
}

func NewOrganizationService(repository OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		repository: repository,
	}
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, organization CreateOrganizationRequest) error {
	if err := s.validateCreateOrganizationRequest(organization); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return s.repository.CreateOrganization(ctx, Organization{
		Name:        organization.Name,
		Logo:        organization.Logo,
		OwnerUserID: organization.OwnerUserID,
	})
}

func (s *OrganizationService) GetOrganizationByID(ctx context.Context, id string) (*Organization, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("organization ID cannot be empty")
	}

	return s.repository.GetOrganizationByID(ctx, id)
}

func (s *OrganizationService) GetOrganizationsByOwner(ctx context.Context, ownerUserID string) ([]Organization, error) {
	if strings.TrimSpace(ownerUserID) == "" {
		return nil, fmt.Errorf("owner user ID cannot be empty")
	}

	return s.repository.GetOrganizationsByOwner(ctx, ownerUserID)
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, id string, organization Organization) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("organization ID cannot be empty")
	}

	if err := s.validateOrganization(organization); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingOrganization, err := s.repository.GetOrganizationByID(ctx, id)
	if err != nil {
		return fmt.Errorf("organization not found: %w", err)
	}

	organization.CreatedAt = existingOrganization.CreatedAt
	organization.UpdatedAt = time.Now()

	return s.repository.UpdateOrganization(ctx, id, organization)
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("organization ID cannot be empty")
	}

	return s.repository.DeleteOrganization(ctx, id)
}

func (s *OrganizationService) ListOrganizations(ctx context.Context, limit, offset int) ([]Organization, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.repository.ListOrganizations(ctx, limit, offset)
}

func (s *OrganizationService) CountOrganizations(ctx context.Context) (int64, error) {
	return s.repository.CountOrganizations(ctx)
}

func (s *OrganizationService) validateOrganization(organization Organization) error {
	if strings.TrimSpace(organization.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(organization.OwnerUserID) == "" {
		return fmt.Errorf("owner user ID is required")
	}

	return nil
}

func (s *OrganizationService) validateCreateOrganizationRequest(organization CreateOrganizationRequest) error {
	if strings.TrimSpace(organization.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(organization.OwnerUserID) == "" {
		return fmt.Errorf("owner user ID is required")
	}

	return nil
}
