package organizations

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBOrganizationRepository struct {
	collection *mongo.Collection
}

func NewMongoDBOrganizationRepository(collection *mongo.Collection) *MongoDBOrganizationRepository {
	return &MongoDBOrganizationRepository{
		collection: collection,
	}
}

func (r *MongoDBOrganizationRepository) CreateOrganization(ctx context.Context, organization Organization) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	now := time.Now()
	organization.CreatedAt = now
	organization.UpdatedAt = now

	if organization.ID == "" {
		organization.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, organization)
	if err != nil {
		return fmt.Errorf("failed to create organization: %w", err)
	}

	return nil
}

func (r *MongoDBOrganizationRepository) GetOrganizationByID(ctx context.Context, id string) (*Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var organization Organization
	log.Println(id)
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&organization)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}

	return &organization, nil
}

func (r *MongoDBOrganizationRepository) GetOrganizationsByOwner(ctx context.Context, ownerUserID string) ([]Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"owner_user_id": ownerUserID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations by owner: %w", err)
	}
	defer cursor.Close(ctx)

	var organizations []Organization
	if err = cursor.All(ctx, &organizations); err != nil {
		return nil, fmt.Errorf("failed to decode organizations: %w", err)
	}

	return organizations, nil
}

func (r *MongoDBOrganizationRepository) UpdateOrganization(ctx context.Context, id string, organization Organization) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	organization.UpdatedAt = time.Now()
	organization.ID = id

	filter := bson.M{"id": id}
	update := bson.M{"$set": organization}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("organization not found")
	}

	return nil
}

func (r *MongoDBOrganizationRepository) DeleteOrganization(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("organization not found")
	}

	return nil
}

func (r *MongoDBOrganizationRepository) ListOrganizations(ctx context.Context, limit, offset int) ([]Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations: %w", err)
	}
	defer cursor.Close(ctx)

	var organizations []Organization
	if err = cursor.All(ctx, &organizations); err != nil {
		return nil, fmt.Errorf("failed to decode organizations: %w", err)
	}

	return organizations, nil
}

func (r *MongoDBOrganizationRepository) CountOrganizations(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("failed to count organizations: %w", err)
	}

	return count, nil
}
