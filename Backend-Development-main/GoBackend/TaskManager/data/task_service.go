package data

import (
	"errors"
	"sync"
	"time"
	"TaskManager/models"
)

// handling task data operations.
type TaskService struct {
	tasks map[int]models.Task
	mu    sync.Mutex
	nextID int
}

//creating a new TaskService.
func NewTaskService() *TaskService {
	s := &TaskService{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}

	// including mock data
	mockTasks := models.MockData()
	for _, task := range mockTasks {
		s.tasks[task.ID] = task
		if task.ID >= s.nextID {
			s.nextID = task.ID + 1
		}
	}

	return s
}

// returning all tasks.
func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// returning a task by ID.
func (s *TaskService) GetTask(id int) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

// creating a new task.
func (s *TaskService) CreateTask(title, description, status string, dueDate time.Time) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := models.Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}
	s.tasks[s.nextID] = task
	s.nextID++
	return task
}

// updating an existing task.
func (s *TaskService) UpdateTask(id int, title, description, status string, dueDate time.Time) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	task.Title = title
	task.Description = description
	task.Status = status
	task.DueDate = dueDate
	s.tasks[id] = task

	return task, nil
}

// deleting a task by ID.
func (s *TaskService) DeleteTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
