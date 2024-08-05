package controllers

import (
	"net/http"
	"time"
	"TaskManager2/data"
	"github.com/gin-gonic/gin"
)

// TaskController handles incoming HTTP requests related to tasks.
type TaskController struct {
	taskService *data.TaskService
}

// NewTaskController creates a new TaskController.
func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

// GetTasks handles GET requests to retrieve all tasks.
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask handles GET requests to retrieve a specific task by ID.
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask handles POST requests to create a new task.
func (tc *TaskController) CreateTask(c *gin.Context) {
	var taskInput struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" binding:"required"`
		DueDate     string `json:"due_date" binding:"required"`
	}

	if err := c.BindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse(time.RFC3339, taskInput.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}

	task, err := tc.taskService.CreateTask(taskInput.Title, taskInput.Description, taskInput.Status, dueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// UpdateTask handles PUT requests to update an existing task by ID.
func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var taskInput struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" binding:"required"`
		DueDate     string `json:"due_date" binding:"required"`
	}

	if err := c.BindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse(time.RFC3339, taskInput.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}

	task, err := tc.taskService.UpdateTask(id, taskInput.Title, taskInput.Description, taskInput.Status, dueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DeleteTask handles DELETE requests to delete a task by ID.
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := tc.taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
