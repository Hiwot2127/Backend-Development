package router

import (
	"TaskManager3/controllers"
	"TaskManager3/data"
	"TaskManager3/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database, secretKey string) *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService(db)
	userService := data.NewUserService(db, secretKey)
	controller := controllers.NewTaskController(taskService, userService, secretKey)

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(secretKey))

	protected.GET("/tasks", controller.GetTasks)
	protected.GET("/tasks/:id", controller.GetTask)
	protected.POST("/tasks", controller.CreateTask)
	protected.PUT("/tasks/:id", controller.UpdateTask)
	protected.DELETE("/tasks/:id", controller.DeleteTask)

	admin := protected.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.GET("/users", controller.ListUsers) 
	}

	return r
}

