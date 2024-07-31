package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	// connect to db
	// todo: load from env
	DbDriver := "postgres"
	DbUser := "postgres"
	DbPassword := "postgres"
	DbPort := "5432"
	DbHost := "localhost"
	DbName := "iam"

	var err error
	DB, err = gorm.Open(DbDriver, "host="+DbHost+" port="+DbPort+" user="+DbUser+" dbname="+DbName+" password="+DbPassword+" sslmode=disable")
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&Role{})
	initRole(DB)
}
