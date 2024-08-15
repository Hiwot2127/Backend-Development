package Usecases

import (
	"TaskManager5/Domain"
	"TaskManager5/Repositories"
)

type UserService struct {
	repo      Repositories.UserRepository
	secretKey string
}

func NewUserService(repo Repositories.UserRepository, secretKey string) *UserService {
	return &UserService{repo: repo, secretKey: secretKey}
}

func (us *UserService) RegisterUser(user Domain.User) (*Domain.User, error) {
	return us.repo.CreateUser(user)
}

func (us *UserService) AuthenticateUser(username, password string) (string, error) {
	return us.repo.AuthenticateUser(username, password)
}

func (us *UserService) GetUserByID(userID string) (*Domain.User, error) {
	return us.repo.GetUserByID(userID)
}

func (us *UserService) GetAllUsers() ([]Domain.User, error) {
	return us.repo.GetAllUsers()
}
