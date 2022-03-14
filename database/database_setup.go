package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// connects the DB
func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Timeout after given time
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx) // connect to the DB
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) //ping to check if connected
	if err != nil {
		fmt.Println("Failed to connect to MongoDB")
		return nil
	}
	fmt.Println("Successfully connected to MongoDB")
	return client
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {}
