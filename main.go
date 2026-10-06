package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-task-api/config"
	"go-task-api/routes"
)

func main() {

	config.ConnectDatabase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GoTask API is running!",
		})
	})

	routes.SetupRoutes(router)

	router.Run(":8080")
}