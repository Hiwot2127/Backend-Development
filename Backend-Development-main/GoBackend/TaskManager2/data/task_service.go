package data

import (
	"context"
	"time"
	"errors"
	"TaskManager2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService provides methods to interact with the tasks collection.
type TaskService struct {
	collection *mongo.Collection
}

// NewTaskService creates a new TaskService.
func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
	}
}

// GetAllTasks retrieves all tasks from the database.
func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTask retrieves a task by ID from the database.
func (s *TaskService) GetTask(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	err = s.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}
	return task, nil
}

// CreateTask creates a new task in the database.
func (s *TaskService) CreateTask(title, description, status string, dueDate time.Time) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}
	_, err := s.collection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// UpdateTask updates an existing task in the database.
func (s *TaskService) UpdateTask(id, title, description, status string, dueDate time.Time) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	update := bson.M{
		"$set": models.Task{
			ID:          objID, // Ensure the ID is set in the update
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			Status:      status,
		},
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return models.Task{}, err
	}
	return s.GetTask(id)
}

// DeleteTask deletes a task by ID from the database.
func (s *TaskService) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
