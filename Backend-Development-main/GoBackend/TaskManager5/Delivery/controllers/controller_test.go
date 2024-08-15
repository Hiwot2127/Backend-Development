package controllers

import (
	"TaskManager5/Domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock TaskService
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskService) GetTask(id primitive.ObjectID) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskService) CreateTask(task Domain.Task) (*Domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(id primitive.ObjectID, updatedTask Domain.Task) (*Domain.Task, error) {
	args := m.Called(id, updatedTask)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskService) DeleteTask(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskService) GetTasksByUserID(userID primitive.ObjectID) ([]Domain.Task, error) {
	args := m.Called(userID)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

// Mock UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) RegisterUser(user Domain.User) (*Domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserService) AuthenticateUser(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) GetUserByID(userID primitive.ObjectID) (*Domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]Domain.User, error) {
	args := m.Called()
	return args.Get(0).([]Domain.User), args.Error(1)
}

// Test GetTasks
func TestTaskController_GetTasks(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	mockTaskService.On("GetTasks").Return([]Domain.Task{
		{ID: primitive.NewObjectID(), Title: "Test Task"},
	}, nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/tasks", tc.GetTasks)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")

	mockTaskService.AssertExpectations(t)
}

// Test GetTask
func TestTaskController_GetTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	taskID := primitive.NewObjectID()
	mockTaskService.On("GetTask", taskID).Return(&Domain.Task{
		ID:    taskID,
		Title: "Test Task",
	}, nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/tasks/:id", tc.GetTask)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")

	mockTaskService.AssertExpectations(t)
}

// Test CreateTask
func TestTaskController_CreateTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	newTask := Domain.Task{Title: "New Task"}
	mockTaskService.On("CreateTask", newTask).Return(&Domain.Task{
		ID:    primitive.NewObjectID(),
		Title: "New Task",
	}, nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("POST", "/tasks", nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/tasks", tc.CreateTask)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Task")

	mockTaskService.AssertExpectations(t)
}

// Test UpdateTask
func TestTaskController_UpdateTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	taskID := primitive.NewObjectID()
	updatedTask := Domain.Task{Title: "Updated Task"}
	mockTaskService.On("UpdateTask", taskID, updatedTask).Return(&Domain.Task{
		ID:    taskID,
		Title: "Updated Task",
	}, nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.PUT("/tasks/:id", tc.UpdateTask)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Task")

	mockTaskService.AssertExpectations(t)
}

// Test DeleteTask
func TestTaskController_DeleteTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	taskID := primitive.NewObjectID()
	mockTaskService.On("DeleteTask", taskID).Return(nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.DELETE("/tasks/:id", tc.DeleteTask)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	mockTaskService.AssertExpectations(t)
}

// Test GetTasksByUserID
func TestTaskController_GetTasksByUserID(t *testing.T) {
	mockTaskService := new(MockTaskService)
	mockUserService := new(MockUserService)

	userID := primitive.NewObjectID()
	mockTaskService.On("GetTasksByUserID", userID).Return([]Domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task for user1"},
	}, nil)

	tc := NewTaskController(mockTaskService, mockUserService, "secretKey")

	req, _ := http.NewRequest("GET", "/tasks/user/"+userID.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/tasks/user/:userID", tc.GetTasksByUserID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task for user1")

	mockTaskService.AssertExpectations(t)
}



