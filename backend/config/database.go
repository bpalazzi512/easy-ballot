package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		URI:      getEnvOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
		Database: getEnvOrDefault("MONGODB_DATABASE", "easy_ballot"),
		Timeout:  10 * time.Second,
	}
}

func ConnectMongoDB(config *DatabaseConfig) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	authSource := getEnvOrDefault("MONGODB_AUTHSOURCE", "admin")
	fmt.Println("username", username)
	fmt.Println("password", password)
	fmt.Println("authSource", authSource)

	clientOptions := options.Client().ApplyURI(config.URI).SetAuth(options.Credential{
		Username:   username,
		Password:   password,
		AuthSource: authSource,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database(config.Database)
	return client, database, nil
}

func CloseMongoDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
