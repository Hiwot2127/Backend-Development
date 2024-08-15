package Usecases

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"TaskManager5/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user Domain.User) (*Domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserRepository) AuthenticateUser(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(userID string) (*Domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]Domain.User, error) {
	args := m.Called()
	return args.Get(0).([]Domain.User), args.Error(1)
}

// Test for RegisterUser
func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo, "secret")

	user := Domain.User{
		ID:       primitive.NewObjectID(),
		Username: "user1",
		Password: "password",
	}
	mockRepo.On("CreateUser", user).Return(&user, nil)

	result, err := service.RegisterUser(user)

	assert.NoError(t, err)
	assert.Equal(t, &user, result)
	mockRepo.AssertExpectations(t)
}

// Test for AuthenticateUser
func TestAuthenticateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo, "secret")

	username := "user1"
	password := "password"
	token := "someJWTToken"
	mockRepo.On("AuthenticateUser", username, password).Return(token, nil)

	result, err := service.AuthenticateUser(username, password)

	assert.NoError(t, err)
	assert.Equal(t, token, result)
	mockRepo.AssertExpectations(t)
}

// Test for GetUserByID
func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo, "secret")

	user := &Domain.User{
		ID:       primitive.NewObjectID(),
		Username: "user1",
		Password: "password",
	}
	mockRepo.On("GetUserByID", user.ID.Hex()).Return(user, nil)

	result, err := service.GetUserByID(user.ID.Hex())

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

// Test for GetAllUsers
func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo, "secret")

	users := []Domain.User{
		{
			ID:       primitive.NewObjectID(),
			Username: "user1",
			Password: "password",
		},
		{
			ID:       primitive.NewObjectID(),
			Username: "user2",
			Password: "password",
		},
	}
	mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockRepo.AssertExpectations(t)
}