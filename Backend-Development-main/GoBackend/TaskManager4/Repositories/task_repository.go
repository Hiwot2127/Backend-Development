package Repositories

import (
	"context"
	"errors"
	"time"

	"TaskManager4/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	GetTasks() ([]Domain.Task, error)
	GetTask(id string) (*Domain.Task, error)
	CreateTask(task Domain.Task) (*Domain.Task, error)
	UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]Domain.Task, error)
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
	}
}

func (tr *taskRepository) GetTasks() ([]Domain.Task, error) {
	var tasks []Domain.Task
	cursor, err := tr.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *taskRepository) GetTask(id string) (*Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	var task Domain.Task
	err = tr.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *taskRepository) CreateTask(task Domain.Task) (*Domain.Task, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := tr.collection.InsertOne(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *taskRepository) UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
			"user_id":     updatedTask.UserID,
			"updated_at":  time.Now(),
		},
	}
	_, err = tr.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return tr.GetTask(id)
}

func (tr *taskRepository) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	_, err = tr.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTasksByUserID(userID string) ([]Domain.Task, error) {
	var tasks []Domain.Task
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := tr.collection.Find(context.Background(), bson.M{"user_id": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
