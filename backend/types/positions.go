package types

import (
	"context"
	"time"
)

type Position struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	OrganizationID string    `json:"organization_id" bson:"organization_id"`
	Title          string    `json:"title" bson:"title"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type CreatePositionRequest struct {
	OrganizationID string `json:"organization_id"`
	Title          string `json:"title"`
}

type PositionRepository interface {
	CreatePosition(ctx context.Context, position Position) error
	GetPositionByID(ctx context.Context, id string) (*Position, error)
	GetPositionsByOrganization(ctx context.Context, organizationID string) ([]Position, error)
	UpdatePosition(ctx context.Context, id string, position Position) error
	DeletePosition(ctx context.Context, id string) error
	ListPositions(ctx context.Context, organizationID string, limit, offset int) ([]Position, error)
	CountPositions(ctx context.Context, organizationID string) (int64, error)
}
