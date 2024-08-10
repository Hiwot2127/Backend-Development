package Usecases


import (
	"TaskManager4/Domain"
	"TaskManager4/Repositories"
)

type TaskService struct {
	repo Repositories.TaskRepository
}

func NewTaskService(repo Repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) GetTasks() ([]Domain.Task, error) {
	return ts.repo.GetTasks()
}

func (ts *TaskService) GetTask(id string) (*Domain.Task, error) {
	return ts.repo.GetTask(id)
}

func (ts *TaskService) CreateTask(task Domain.Task) (*Domain.Task, error) {
	return ts.repo.CreateTask(task)
}

func (ts *TaskService) UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error) {
	return ts.repo.UpdateTask(id, updatedTask)
}

func (ts *TaskService) DeleteTask(id string) error {
	return ts.repo.DeleteTask(id)
}

func (ts *TaskService) GetTasksByUserID(userID string) ([]Domain.Task, error) {
	return ts.repo.GetTasksByUserID(userID)
}
