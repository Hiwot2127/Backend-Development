package routers

import (
	"TaskManager5/Delivery/controllers"
	"TaskManager5/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, controller *controllers.TaskController, secretKey string) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Authenticated routes
	r.Use(Infrastructure.AuthMiddleware(secretKey))

	// Task routes
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTask)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)

	// Admin routes
	r.Use(Infrastructure.AdminMiddleware())
	r.GET("/admin/users", controller.ListUsers)
	r.GET("/admin/users/:id", controller.GetUserByID)
	r.GET("/admin/tasks/user/:user_id", controller.GetTasksByUserID)
}
