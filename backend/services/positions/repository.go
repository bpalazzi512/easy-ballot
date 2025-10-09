package positions

import (
	"context"
	"fmt"
	"time"

	"github.com/bpalazzi512/easy-ballot/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBPositionRepository struct {
	collection *mongo.Collection
}

func NewMongoDBPositionRepository(collection *mongo.Collection) *MongoDBPositionRepository {
	return &MongoDBPositionRepository{
		collection: collection,
	}
}

func (r *MongoDBPositionRepository) CreatePosition(ctx context.Context, position types.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	now := time.Now()
	position.CreatedAt = now
	position.UpdatedAt = now

	if position.ID == "" {
		position.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, position)
	if err != nil {
		return fmt.Errorf("failed to create position: %w", err)
	}

	return nil
}

func (r *MongoDBPositionRepository) GetPositionByID(ctx context.Context, id string) (*types.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var position types.Position
	filter := bson.M{"id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&position)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("position not found")
		}
		return nil, fmt.Errorf("failed to get position: %w", err)
	}

	return &position, nil
}

func (r *MongoDBPositionRepository) GetPositionsByOrganization(ctx context.Context, organizationID string) ([]types.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"organization_id": organizationID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get positions by organization: %w", err)
	}
	defer cursor.Close(ctx)

	var positions []types.Position
	if err = cursor.All(ctx, &positions); err != nil {
		return nil, fmt.Errorf("failed to decode positions: %w", err)
	}

	return positions, nil
}

func (r *MongoDBPositionRepository) UpdatePosition(ctx context.Context, id string, position types.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	position.UpdatedAt = time.Now()
	position.ID = id

	filter := bson.M{"id": id}
	update := bson.M{"$set": position}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update position: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("position not found")
	}

	return nil
}

func (r *MongoDBPositionRepository) DeletePosition(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete position: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("position not found")
	}

	return nil
}

func (r *MongoDBPositionRepository) ListPositions(ctx context.Context, organizationID string, limit, offset int) ([]types.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if organizationID != "" {
		filter["organization_id"] = organizationID
	}

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list positions: %w", err)
	}
	defer cursor.Close(ctx)

	var positions []types.Position
	if err = cursor.All(ctx, &positions); err != nil {
		return nil, fmt.Errorf("failed to decode positions: %w", err)
	}

	return positions, nil
}

func (r *MongoDBPositionRepository) CountPositions(ctx context.Context, organizationID string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if organizationID != "" {
		filter["organization_id"] = organizationID
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count positions: %w", err)
	}

	return count, nil
}
