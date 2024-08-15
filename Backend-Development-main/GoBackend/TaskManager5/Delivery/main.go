package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"TaskManager5/Delivery/controllers"
	"TaskManager5/Delivery/router"
	"TaskManager5/Repositories"
	"TaskManager5/Usecases"
)

const (
	mongoURI  = "mongodb://localhost:27017"
	dbName    = "task_manager"
	secretKey = "s5e8ydy9GrJwXJf5cF6Sb58y4KpIhR9+Z1kqfO3D6R0="
)

func main() {
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(dbName)

	taskRepo := Repositories.NewTaskRepository(db)
	userRepo := Repositories.NewUserRepository(db, secretKey)

	taskService := Usecases.NewTaskService(taskRepo)
	userService := Usecases.NewUserService(userRepo, secretKey)

	controller := controllers.NewTaskController(taskService, userService, secretKey)

	r := gin.Default()
	routers.SetupRoutes(r, controller, secretKey)


	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
