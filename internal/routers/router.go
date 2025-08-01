package routers

import (
	"learning-french-service/internal/controller"
	"learning-french-service/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())

	// Add basic health check endpoint
	router.GET("/health", heatlhCheck)

	// API v1 group
	v1 := router.Group("/api/v1")

	v1.GET("/users", controller.NewUserController().GetUsers)

	return router
}

func heatlhCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "okkkk",
		"message": "Apprendre.ai API is running",
		"version": "1.0.0",
	})
}
