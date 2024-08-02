package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
)

//var DB *gorm.DB

type DataBase struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}

var DB DataBase

func ConnectDB() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	/**
	 * Connect to Postgres
	 */
	DbDriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbPort := os.Getenv("DB_PORT")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DATABASE")

	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	DB.Postgres, err = gorm.Open(DbDriver, connectionUrl)
	if err != nil {
		panic(err)
	}

	/**
	 * Connect to Redis
	 */
	RdPort := os.Getenv("REDIS_PORT")
	RdHost := os.Getenv("REDIS_HOST")
	RdDatabase, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	RdPassword := os.Getenv("REDIS_PASSWORD")

	DB.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RdHost, RdPort),
		Password: RdPassword,
		DB:       RdDatabase,
	})

	/**
	 * Migration data
	 */
	DB.Postgres.AutoMigrate(&Account{})
	DB.Postgres.AutoMigrate(&Role{})
	InitRole(DB.Postgres)
}
