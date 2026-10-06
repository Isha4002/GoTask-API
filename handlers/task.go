package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-task-api/config"
	"go-task-api/models"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Create Task
func CreateTask(c *gin.Context) {

	var request CreateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	task := models.Task{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
		UserID:      uint(userID.(float64)),
	}

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

// Get all tasks for logged-in user
func GetTasks(c *gin.Context) {

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var tasks []models.Task

	if err := config.DB.
		Where("user_id = ?", uint(userID.(float64))).
		Find(&tasks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

// Get single task by ID
func GetTask(c *gin.Context) {

	taskID := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var task models.Task

	// Find task only if it belongs to the logged-in user
	if err := config.DB.
		Where("id = ? AND user_id = ?", taskID, uint(userID.(float64))).
		First(&task).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// Update Task
func UpdateTask(c *gin.Context) {

	taskID := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var task models.Task

	// Find task belonging to the logged-in user
	if err := config.DB.
		Where("id = ? AND user_id = ?", taskID, uint(userID.(float64))).
		First(&task).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	var request UpdateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	// Update task
	task.Title = request.Title
	task.Description = request.Description
	task.Completed = request.Completed

	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}

// Delete Task
func DeleteTask(c *gin.Context) {

	taskID := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var task models.Task

	// Find task belonging to logged-in user
	if err := config.DB.
		Where("id = ? AND user_id = ?", taskID, uint(userID.(float64))).
		First(&task).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	// Delete task
	if err := config.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}