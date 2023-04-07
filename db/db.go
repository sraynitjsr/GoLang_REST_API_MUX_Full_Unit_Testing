package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

type Database interface {
	Connect() error
	Disconnect()
	GetCollection(collectionName string) *mongo.Collection
}

type MongoDB struct {
	URI      string
	Database string
}

func (m *MongoDB) Connect() error {
	// Set client options
	clientOptions := options.Client().ApplyURI(m.URI)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	// Set the collection
	collection = client.Database(m.Database).Collection("students")

	return nil
}

func (m *MongoDB) Disconnect() {
	// Disconnect from MongoDB
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Error disconnecting from MongoDB: ", err)
	}

	log.Println("Disconnected from MongoDB!")
}

func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return client.Database(m.Database).Collection(collectionName)
}
