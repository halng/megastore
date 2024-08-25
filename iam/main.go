package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/handlers"
	"github.com/tanhaok/megastore/kafka"
	"github.com/tanhaok/megastore/logging"
	"github.com/tanhaok/megastore/models"
	"os"
)

func main() {
	logging.InitLogging()

	// connect database
	db.ConnectDB()
	models.Initialize()

	var err error

	_ = godotenv.Load(".env")

	// init kafka server
	bootstrapServer := os.Getenv("KAFKA_HOST")
	err = kafka.InitializeKafkaProducer(bootstrapServer)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	router := gin.Default()

	groupV1 := router.Group("/api/v1")

	// routes
	groupV1.POST("/login", handlers.Login)
	groupV1.POST("/register", handlers.Register)
	groupV1.GET("/validate", handlers.Validate)

	err = router.Run(":" + port)
	logging.LOGGER.Info(fmt.Sprintf("Starting web service on port %s", port))
	if err != nil {
		return
	}
}
