package Usecases

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"TaskManager5/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock repository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTask(id string) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) CreateTask(task Domain.Task) (*Domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error) {
	args := m.Called(id, updatedTask)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasksByUserID(userID string) ([]Domain.Task, error) {
	args := m.Called(userID)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

// Test for GetTasks
func TestGetTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	tasks := []Domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now(),
			Status:      "Completed",
			UserID:      primitive.NewObjectID(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	mockRepo.On("GetTasks").Return(tasks, nil)

	result, err := service.GetTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

// Test for GetTask
func TestGetTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	task := &Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "Pending",
		UserID:      primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("GetTask", task.ID.Hex()).Return(task, nil)

	result, err := service.GetTask(task.ID.Hex())

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

// Test for CreateTask
func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	task := Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "New Task",
		Description: "New Task Description",
		DueDate:     time.Now(),
		Status:      "Pending",
		UserID:      primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("CreateTask", task).Return(&task, nil)

	result, err := service.CreateTask(task)

	assert.NoError(t, err)
	assert.Equal(t, &task, result)
	mockRepo.AssertExpectations(t)
}

// Test for UpdateTask
func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	updatedTask := Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "In Progress",
		UserID:      primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("UpdateTask", updatedTask.ID.Hex(), updatedTask).Return(&updatedTask, nil)

	result, err := service.UpdateTask(updatedTask.ID.Hex(), updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, &updatedTask, result)
	mockRepo.AssertExpectations(t)
}

// Test for DeleteTask
func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	taskID := primitive.NewObjectID().Hex()
	mockRepo.On("DeleteTask", taskID).Return(nil)

	err := service.DeleteTask(taskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test for GetTasksByUserID
func TestGetTasksByUserID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	tasks := []Domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
			Status:      "Pending",
			UserID:      primitive.NewObjectID(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	mockRepo.On("GetTasksByUserID", "user1").Return(tasks, nil)

	result, err := service.GetTasksByUserID("user1")

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

