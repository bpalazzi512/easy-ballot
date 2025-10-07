package organizations

import (
	"context"
	"time"
)

type Organization struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Logo        string    `json:"logo" bson:"logo"`
	OwnerUserID string    `json:"owner_user_id" bson:"owner_user_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateOrganizationRequest struct {
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	OwnerUserID string `json:"owner_user_id"`
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, organization Organization) error
	GetOrganizationByID(ctx context.Context, id string) (*Organization, error)
	GetOrganizationsByOwner(ctx context.Context, ownerUserID string) ([]Organization, error)
	UpdateOrganization(ctx context.Context, id string, organization Organization) error
	DeleteOrganization(ctx context.Context, id string) error
	ListOrganizations(ctx context.Context, limit, offset int) ([]Organization, error)
	CountOrganizations(ctx context.Context) (int64, error)
}
