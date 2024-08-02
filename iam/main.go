package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tanhaok/MyStore/handlers"
	"github.com/tanhaok/MyStore/models"
	"log"
	"os"
)

func main() {

	// connect database
	models.ConnectDB()

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := gin.Default()

	groupV1 := router.Group("/api/v1")

	// routes
	groupV1.POST("/login", handlers.Login)
	groupV1.POST("/register", handlers.Register)
	groupV1.GET("/validate", handlers.Validate)

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
