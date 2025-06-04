package routes

import (
	"github.com/yourname/go-task-tracker/controllers"
	"github.com/gin-gonic/gin"
	"github.com/yourname/go-task-tracker/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)

	server.Use(middlewares.Authenticate)
	server.POST("/project", controllers.CreateProject)
	server.DELETE("/project/:id", controllers.DeleteProject)
	server.GET("/project/:id", controllers.GetProjectById)
	server.PUT("/project/:id", controllers.UpdateProject)

	server.POST("/task", controllers.CreateTask)
	server.GET("/task/:id", controllers.GetTaskById)
	server.GET("/task/project/:id", controllers.GetTasksByProjectId)
	server.PUT("/task/:id", controllers.UpdateTask)
	server.DELETE("/task/:id", controllers.DeleteTask)
}