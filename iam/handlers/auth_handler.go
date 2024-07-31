package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/MyStore/constants"
	"github.com/tanhaok/MyStore/models"
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

// ========================= Functions =========================

// Register TODO: check if username or email already exists in database
// Register new account
func Register(c *gin.Context) {
	var userInput RegisterRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.MessageError})
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

	c.JSON(http.StatusCreated, gin.H{"msg": constants.AccountCreated})
}

// Login verify user credentials and return uuid pair with token saved in redis
func Login(c *gin.Context) {
	// login{

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// validate

}
