package controllers

import (
	"net/http"
	"strconv"
	"TaskManager/data"
	"TaskManager/models"

	"github.com/gin-gonic/gin"
)

// handling HTTP requests for tasks.
type TaskController struct {
	service *data.TaskService
}

// creating a new TaskController.
func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{service: service}
}

// handling GET /tasks.
func (ctl *TaskController) GetTasks(c *gin.Context) {
	tasks := ctl.service.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// handling GET /tasks/:id.
func (ctl *TaskController) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	task, err := ctl.service.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// handling POST /tasks.
func (ctl *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task = ctl.service.CreateTask(task.Title, task.Description, task.Status, task.DueDate)
	c.JSON(http.StatusCreated, task)
}

// handling PUT /tasks/:id.
func (ctl *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := ctl.service.UpdateTask(id, task.Title, task.Description, task.Status, task.DueDate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// handling DELETE /tasks/:id.
func (ctl *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := ctl.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task has been deleted"})
}
