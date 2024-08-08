package main

import (
	"context"
	"log"
	"time"

	"TaskManager3/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI  = "mongodb://localhost:27017"
	dbName    = "task_manager"
	secretKey = "s5e8ydy9GrJwXJf5cF6Sb58y4KpIhR9+Z1kqfO3D6R0=" 
)

func main() {
	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect the client when the main function finishes
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(dbName)

	// Setup router
	r := router.SetupRouter(db, secretKey)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

