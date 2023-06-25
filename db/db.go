package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

var Collection *mongo.Collection
var UserCollection *mongo.Collection
var Ctx = context.TODO()

func InitDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("task_manager")
	Collection = database.Collection("tasks")
	UserCollection = database.Collection("users")

	// Create the "tasks" collection if it doesn't exist
	if err := createCollection(database, "tasks"); err != nil {
		log.Fatal(err)
	}

	if err := createCollection(database, "users"); err != nil {
		log.Fatal(err)
	}
}

func createCollection(database *mongo.Database, collectionName string) error {
	collectionOptions := options.CreateCollection()
	err := database.CreateCollection(Ctx, collectionName, collectionOptions)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return fmt.Errorf("failed to create collection: %v", err)
	}
	return nil
}
