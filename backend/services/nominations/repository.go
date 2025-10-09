package nominations

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

type MongoDBNominationRepository struct {
	collection *mongo.Collection
}

func NewMongoDBNominationRepository(collection *mongo.Collection) *MongoDBNominationRepository {
	return &MongoDBNominationRepository{
		collection: collection,
	}
}

func (r *MongoDBNominationRepository) CreateNomination(ctx context.Context, nomination types.Nomination) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	now := time.Now()
	nomination.CreatedAt = now
	nomination.UpdatedAt = now
	nomination.Status = types.NominationStatusPending

	if nomination.ID == "" {
		nomination.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, nomination)
	if err != nil {
		return fmt.Errorf("failed to create nomination: %w", err)
	}

	return nil
}

func (r *MongoDBNominationRepository) GetNominationByID(ctx context.Context, id string) (*types.Nomination, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var nomination types.Nomination
	filter := bson.M{"id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&nomination)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("nomination not found")
		}
		return nil, fmt.Errorf("failed to get nomination: %w", err)
	}

	return &nomination, nil
}

func (r *MongoDBNominationRepository) GetNominationsByPosition(ctx context.Context, positionID string) ([]types.Nomination, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"position_id": positionID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get nominations by position: %w", err)
	}
	defer cursor.Close(ctx)

	var nominations []types.Nomination
	if err = cursor.All(ctx, &nominations); err != nil {
		return nil, fmt.Errorf("failed to decode nominations: %w", err)
	}

	return nominations, nil
}

func (r *MongoDBNominationRepository) GetNominationsByNominee(ctx context.Context, nomineeID string) ([]types.Nomination, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"nominee_id": nomineeID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get nominations by nominee: %w", err)
	}
	defer cursor.Close(ctx)

	var nominations []types.Nomination
	if err = cursor.All(ctx, &nominations); err != nil {
		return nil, fmt.Errorf("failed to decode nominations: %w", err)
	}

	return nominations, nil
}

func (r *MongoDBNominationRepository) GetNominationsByNominator(ctx context.Context, nominatorID string) ([]types.Nomination, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"nominator_id": nominatorID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get nominations by nominator: %w", err)
	}
	defer cursor.Close(ctx)

	var nominations []types.Nomination
	if err = cursor.All(ctx, &nominations); err != nil {
		return nil, fmt.Errorf("failed to decode nominations: %w", err)
	}

	return nominations, nil
}

func (r *MongoDBNominationRepository) UpdateNomination(ctx context.Context, id string, nomination types.Nomination) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	nomination.UpdatedAt = time.Now()
	nomination.ID = id

	filter := bson.M{"id": id}
	update := bson.M{"$set": nomination}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update nomination: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("nomination not found")
	}

	return nil
}

func (r *MongoDBNominationRepository) UpdateNominationStatus(ctx context.Context, id string, status types.NominationStatus) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update nomination status: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("nomination not found")
	}

	return nil
}

func (r *MongoDBNominationRepository) DeleteNomination(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete nomination: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("nomination not found")
	}

	return nil
}

func (r *MongoDBNominationRepository) ListNominations(ctx context.Context, positionID string, limit, offset int) ([]types.Nomination, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if positionID != "" {
		filter["position_id"] = positionID
	}

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list nominations: %w", err)
	}
	defer cursor.Close(ctx)

	var nominations []types.Nomination
	if err = cursor.All(ctx, &nominations); err != nil {
		return nil, fmt.Errorf("failed to decode nominations: %w", err)
	}

	return nominations, nil
}

func (r *MongoDBNominationRepository) CountNominations(ctx context.Context, positionID string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if positionID != "" {
		filter["position_id"] = positionID
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count nominations: %w", err)
	}

	return count, nil
}
