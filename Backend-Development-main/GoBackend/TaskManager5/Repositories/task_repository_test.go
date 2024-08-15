package Repositories

import (
	"context"
	"testing"
	"time"

	"TaskManager5/Domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection is a mock implementation of the mongo.Collection interface
type MockCollection struct {
	mock.Mock
}

// Mock InsertOne method
func (m *MockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

// Mock Find method
func (m *MockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

// Mock FindOne method
func (m *MockCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

// Mock UpdateOne method
func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

// Mock DeleteOne method
func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

// TestTaskRepository tests the taskRepository methods
func TestTaskRepository(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &taskRepository{collection: mockCollection}

	// Test CreateTask
	t.Run("CreateTask", func(t *testing.T) {
		task := Domain.Task{
			Title:       "Test Task",
			Description: "This is a test task",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
		}
		mockCollection.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).
			Return(&mongo.InsertOneResult{InsertedID: task.ID}, nil)

		createdTask, err := repo.CreateTask(task)
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)
		assert.Equal(t, task.Title, createdTask.Title)
		assert.Equal(t, task.Description, createdTask.Description)

		mockCollection.AssertExpectations(t)
	})

	// Test GetTasks
	t.Run("GetTasks", func(t *testing.T) {
		mockCursor := new(mongo.Cursor)
		mockCollection.On("Find", mock.Anything, mock.Anything, mock.Anything).
			Return(mockCursor, nil)

		tasks, err := repo.GetTasks()
		assert.NoError(t, err)
		assert.NotNil(t, tasks)

		mockCollection.AssertExpectations(t)
	})

	// Test GetTask
	t.Run("GetTask", func(t *testing.T) {
		task := Domain.Task{
			ID:          primitive.NewObjectID(),
			Title:       "Unique Task",
			Description: "Unique Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
		}
		mockSingleResult := mongo.NewSingleResultFromDocument(task, nil, nil)
		mockCollection.On("FindOne", mock.Anything, mock.Anything, mock.Anything).
			Return(mockSingleResult)

		fetchedTask, err := repo.GetTask(task.ID.Hex())
		assert.NoError(t, err)
		assert.NotNil(t, fetchedTask)
		assert.Equal(t, task.Title, fetchedTask.Title)

		mockCollection.AssertExpectations(t)
	})

	// Test UpdateTask
	t.Run("UpdateTask", func(t *testing.T) {
		task := Domain.Task{
			ID:          primitive.NewObjectID(),
			Title:       "Task to Update",
			Description: "Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
		}
		updatedTask := Domain.Task{
			Title:       "Updated Title",
			Description: "Updated Description",
			DueDate:     time.Now().Add(48 * time.Hour),
			Status:      "Completed",
			UserID:      task.UserID,
		}
		mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&mongo.UpdateResult{}, nil)
		mockCollection.On("FindOne", mock.Anything, mock.Anything, mock.Anything).
			Return(mongo.NewSingleResultFromDocument(updatedTask, nil, nil))

		result, err := repo.UpdateTask(task.ID.Hex(), updatedTask)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedTask.Title, result.Title)

		mockCollection.AssertExpectations(t)
	})

	// Test DeleteTask
	t.Run("DeleteTask", func(t *testing.T) {
		task := Domain.Task{
			ID:          primitive.NewObjectID(),
			Title:       "Task to Delete",
			Description: "Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
		}
		mockCollection.On("DeleteOne", mock.Anything, mock.Anything, mock.Anything).
			Return(&mongo.DeleteResult{}, nil)

		err := repo.DeleteTask(task.ID.Hex())
		assert.NoError(t, err)

		mockCollection.AssertExpectations(t)
	})
}

