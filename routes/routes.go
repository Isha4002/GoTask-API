package routes

import (
	"github.com/gin-gonic/gin"

	"go-task-api/handlers"
	"go-task-api/middleware"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	// Public routes
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	// Protected routes
	protected := api.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("", handlers.CreateTask)
		protected.GET("", handlers.GetTasks)
		protected.GET("/:id", handlers.GetTask)
		protected.PUT("/:id", handlers.UpdateTask)
		protected.DELETE("/:id", handlers.DeleteTask)
	}
}