package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/MyStore/constants"
	"github.com/tanhaok/MyStore/models"
	"log"
	"net/http"
)

// ========================= Structs =========================

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ========================= Functions =========================

// Register new account
func Register(c *gin.Context) {
	var userInput RegisterRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.MessageError})
		return
	}


	if _, err := models.ExistsByEmailOrUsername(userInput.Email, userInput.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf(constants.AccountExists, userInput.Email, userInput.Username)})
		return
	}

	account := models.Account{}
	account.Email = userInput.Email
	account.Username = userInput.Username
	account.Password = userInput.Password
	account.LastName = userInput.LastName
	account.FirstName = userInput.FirstName
	_, err := account.SaveAccount()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.MessageError})
		return
	}

	// send msg to kafka

	c.JSON(http.StatusCreated, gin.H{"msg": constants.AccountCreated})
}

// Login verify user credentials and return uuid pair with token saved in redis
func Login(c *gin.Context) {
	var userInput LoginRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Printf("Value is incorrect. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.MessageError})
		return
	}

	var account models.Account
	var err error

	if account, err = models.ExistsByEmailOrUsername(userInput.Email, userInput.Username); err != nil {
		log.Printf("Account donesn't exits")
		c.JSON(http.StatusNotFound, gin.H{"msg": constants.AccountNotFound})
		return
	}

	if !account.ComparePassword(userInput.Password) {
		log.Print("Password doesn't match")
		c.JSON(http.StatusUnauthorized, gin.H{"msg": constants.Unauthorized})
		return
	}

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// validate

}
