package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient interface {
	DB() (*mongo.Database, error)
	Close() error
}

type mongoDBClient struct {
	client *mongo.Client
}

func NewMongoDBConnection() (MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbURL, err := getURLLocal()
	if err != nil {
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &mongoDBClient{client: client}, nil
}

func getURLLocal() (string, error) {
	return os.Getenv("DATABASE_URL"), nil
}

func (c *mongoDBClient) DB() (*mongo.Database, error) {
	return c.client.Database(os.Getenv("DATABASE_NAME")), nil
}

func (c *mongoDBClient) Close() error {
	ctx := context.TODO()
	err := c.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
