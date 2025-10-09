package nominations

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bpalazzi512/easy-ballot/backend/types"
)

type NominationRepository = types.NominationRepository

type NominationService struct {
	repository NominationRepository
}

func NewNominationService(repository NominationRepository) *NominationService {
	return &NominationService{
		repository: repository,
	}
}

func (s *NominationService) CreateNomination(ctx context.Context, nomination types.CreateNominationRequest) error {
	if err := s.validateCreateNominationRequest(nomination); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return s.repository.CreateNomination(ctx, types.Nomination{
		PositionID:  nomination.PositionID,
		NomineeID:   nomination.NomineeID,
		NominatorID: nomination.NominatorID,
		Status:      types.NominationStatusPending,
	})
}

func (s *NominationService) GetNominationByID(ctx context.Context, id string) (*types.Nomination, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("nomination ID cannot be empty")
	}

	return s.repository.GetNominationByID(ctx, id)
}

func (s *NominationService) GetNominationsByPosition(ctx context.Context, positionID string) ([]types.Nomination, error) {
	if strings.TrimSpace(positionID) == "" {
		return nil, fmt.Errorf("position ID cannot be empty")
	}

	return s.repository.GetNominationsByPosition(ctx, positionID)
}

func (s *NominationService) GetNominationsByNominee(ctx context.Context, nomineeID string) ([]types.Nomination, error) {
	if strings.TrimSpace(nomineeID) == "" {
		return nil, fmt.Errorf("nominee ID cannot be empty")
	}

	return s.repository.GetNominationsByNominee(ctx, nomineeID)
}

func (s *NominationService) GetNominationsByNominator(ctx context.Context, nominatorID string) ([]types.Nomination, error) {
	if strings.TrimSpace(nominatorID) == "" {
		return nil, fmt.Errorf("nominator ID cannot be empty")
	}

	return s.repository.GetNominationsByNominator(ctx, nominatorID)
}

func (s *NominationService) UpdateNomination(ctx context.Context, id string, nomination types.Nomination) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("nomination ID cannot be empty")
	}

	if err := s.validateNomination(nomination); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingNomination, err := s.repository.GetNominationByID(ctx, id)
	if err != nil {
		return fmt.Errorf("nomination not found: %w", err)
	}

	nomination.CreatedAt = existingNomination.CreatedAt
	nomination.UpdatedAt = time.Now()

	return s.repository.UpdateNomination(ctx, id, nomination)
}

func (s *NominationService) UpdateNominationStatus(ctx context.Context, id string, status types.NominationStatus) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("nomination ID cannot be empty")
	}

	if err := s.validateNominationStatus(status); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return s.repository.UpdateNominationStatus(ctx, id, status)
}

func (s *NominationService) DeleteNomination(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("nomination ID cannot be empty")
	}

	return s.repository.DeleteNomination(ctx, id)
}

func (s *NominationService) ListNominations(ctx context.Context, positionID string, limit, offset int) ([]types.Nomination, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.repository.ListNominations(ctx, positionID, limit, offset)
}

func (s *NominationService) CountNominations(ctx context.Context, positionID string) (int64, error) {
	return s.repository.CountNominations(ctx, positionID)
}

func (s *NominationService) validateNomination(nomination types.Nomination) error {
	if strings.TrimSpace(nomination.PositionID) == "" {
		return fmt.Errorf("position ID is required")
	}
	if strings.TrimSpace(nomination.NomineeID) == "" {
		return fmt.Errorf("nominee ID is required")
	}
	if strings.TrimSpace(nomination.NominatorID) == "" {
		return fmt.Errorf("nominator ID is required")
	}
	if err := s.validateNominationStatus(nomination.Status); err != nil {
		return err
	}

	return nil
}

func (s *NominationService) validateCreateNominationRequest(nomination types.CreateNominationRequest) error {
	if strings.TrimSpace(nomination.PositionID) == "" {
		return fmt.Errorf("position ID is required")
	}
	if strings.TrimSpace(nomination.NomineeID) == "" {
		return fmt.Errorf("nominee ID is required")
	}
	if strings.TrimSpace(nomination.NominatorID) == "" {
		return fmt.Errorf("nominator ID is required")
	}

	return nil
}

func (s *NominationService) validateNominationStatus(status types.NominationStatus) error {
	switch status {
	case types.NominationStatusAccepted, types.NominationStatusDeclined, types.NominationStatusPending:
		return nil
	default:
		return fmt.Errorf("invalid nomination status: %s", status)
	}
}
