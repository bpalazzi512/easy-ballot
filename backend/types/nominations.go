package types

import (
	"context"
	"time"
)

type NominationStatus string

const (
	NominationStatusAccepted NominationStatus = "accepted"
	NominationStatusDeclined NominationStatus = "declined"
	NominationStatusPending  NominationStatus = "pending"
)

type Nomination struct {
	ID          string           `json:"id" bson:"_id,omitempty"`
	PositionID  string           `json:"position_id" bson:"position_id"`
	NomineeID   string           `json:"nominee_id" bson:"nominee_id"`
	NominatorID string           `json:"nominator_id" bson:"nominator_id"`
	Status      NominationStatus `json:"status" bson:"status"`
	CreatedAt   time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" bson:"updated_at"`
}

type CreateNominationRequest struct {
	PositionID  string `json:"position_id"`
	NomineeID   string `json:"nominee_id"`
	NominatorID string `json:"nominator_id"`
}

type UpdateNominationStatusRequest struct {
	Status NominationStatus `json:"status"`
}

type NominationRepository interface {
	CreateNomination(ctx context.Context, nomination Nomination) error
	GetNominationByID(ctx context.Context, id string) (*Nomination, error)
	GetNominationsByPosition(ctx context.Context, positionID string) ([]Nomination, error)
	GetNominationsByNominee(ctx context.Context, nomineeID string) ([]Nomination, error)
	GetNominationsByNominator(ctx context.Context, nominatorID string) ([]Nomination, error)
	UpdateNomination(ctx context.Context, id string, nomination Nomination) error
	UpdateNominationStatus(ctx context.Context, id string, status NominationStatus) error
	DeleteNomination(ctx context.Context, id string) error
	ListNominations(ctx context.Context, positionID string, limit, offset int) ([]Nomination, error)
	CountNominations(ctx context.Context, positionID string) (int64, error)
}
