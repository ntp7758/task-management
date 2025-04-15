package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	log.Println("Connected to MongoDB")
	return client
}

func GetCollection(dbName, collName string) *mongo.Collection {
	return client.Database(dbName).Collection(collName)
}
