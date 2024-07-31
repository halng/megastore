package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/MyStore/handlers"
	"github.com/tanhaok/MyStore/models"
	"os"
)

func init() {
	//	initialize

}

func main() {

	// connect database
	models.ConnectDB()

	// get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "5051"
	}

	router := gin.Default()

	groupV1 := router.Group("/api/v1")

	// routes
	groupV1.POST("/login", handlers.Login)
	groupV1.POST("/register", handlers.Register)
	groupV1.GET("/validate", handlers.Validate)

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
