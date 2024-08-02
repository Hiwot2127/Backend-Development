package router

import (
	"TaskManager/controllers"
	"TaskManager/data"

	"github.com/gin-gonic/gin"
)

// initializing the routes.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTask)
	r.POST("/tasks", taskController.CreateTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
