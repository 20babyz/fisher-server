package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

var Client *mongo.Client

func InitDB(user, password, host, dbname string) error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?authSource=admin", user, password, host, dbname)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	Client = client
	fmt.Println("Connected to MongoDB successfully")
	return nil
}
