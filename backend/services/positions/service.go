package positions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bpalazzi512/easy-ballot/backend/types"
)

type PositionRepository = types.PositionRepository

type PositionService struct {
	repository PositionRepository
}

func NewPositionService(repository PositionRepository) *PositionService {
	return &PositionService{
		repository: repository,
	}
}

func (s *PositionService) CreatePosition(ctx context.Context, position types.CreatePositionRequest) error {
	if err := s.validateCreatePositionRequest(position); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return s.repository.CreatePosition(ctx, types.Position{
		OrganizationID: position.OrganizationID,
		Title:          position.Title,
	})
}

func (s *PositionService) GetPositionByID(ctx context.Context, id string) (*types.Position, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("position ID cannot be empty")
	}

	return s.repository.GetPositionByID(ctx, id)
}

func (s *PositionService) GetPositionsByOrganization(ctx context.Context, organizationID string) ([]types.Position, error) {
	if strings.TrimSpace(organizationID) == "" {
		return nil, fmt.Errorf("organization ID cannot be empty")
	}

	return s.repository.GetPositionsByOrganization(ctx, organizationID)
}

func (s *PositionService) UpdatePosition(ctx context.Context, id string, position types.Position) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("position ID cannot be empty")
	}

	if err := s.validatePosition(position); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingPosition, err := s.repository.GetPositionByID(ctx, id)
	if err != nil {
		return fmt.Errorf("position not found: %w", err)
	}

	position.CreatedAt = existingPosition.CreatedAt
	position.UpdatedAt = time.Now()

	return s.repository.UpdatePosition(ctx, id, position)
}

func (s *PositionService) DeletePosition(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("position ID cannot be empty")
	}

	return s.repository.DeletePosition(ctx, id)
}

func (s *PositionService) ListPositions(ctx context.Context, organizationID string, limit, offset int) ([]types.Position, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.repository.ListPositions(ctx, organizationID, limit, offset)
}

func (s *PositionService) CountPositions(ctx context.Context, organizationID string) (int64, error) {
	return s.repository.CountPositions(ctx, organizationID)
}

func (s *PositionService) validatePosition(position types.Position) error {
	if strings.TrimSpace(position.OrganizationID) == "" {
		return fmt.Errorf("organization ID is required")
	}
	if strings.TrimSpace(position.Title) == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}

func (s *PositionService) validateCreatePositionRequest(position types.CreatePositionRequest) error {
	if strings.TrimSpace(position.OrganizationID) == "" {
		return fmt.Errorf("organization ID is required")
	}
	if strings.TrimSpace(position.Title) == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}
