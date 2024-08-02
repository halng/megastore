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

type LoginResponse struct {
	ApiToken string `json:"api-token"`
}

// ========================= Functions =========================

// Register new account
func Register(c *gin.Context) {
	var userInput RegisterRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, GetResponseDTO(400, nil, ErrorDTO{constants.MessageError}))
		return
	}

	if models.ExistsByEmailOrUsername(userInput.Email, userInput.Username) {
		c.JSON(http.StatusBadRequest, GetResponseDTO(400, nil, ErrorDTO{fmt.Sprintf(constants.AccountExists, userInput.Email, userInput.Username)}))
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
		c.JSON(http.StatusBadRequest, GetResponseDTO(400, nil, ErrorDTO{constants.MessageError}))
		return
	}

	// send msg to kafka

	c.JSON(http.StatusCreated, GetResponseDTO(201, nil, ErrorDTO{}))
}

// Login verify user credentials and return uuid pair with token saved in redis
func Login(c *gin.Context) {
	var userInput LoginRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Printf("Value is incorrect. %v", err)
		c.JSON(http.StatusBadRequest, GetResponseDTO(404, nil, ErrorDTO{constants.MessageError}))
		return
	}

	var account models.Account
	var err error

	if account, err = models.GetAccountByEmailOrUsername(userInput.Email, userInput.Username); err != nil {
		log.Printf("Account donesn't exits")
		c.JSON(http.StatusNotFound, GetResponseDTO(404, nil, ErrorDTO{constants.AccountNotFound}))
		return
	}

	if !account.ComparePassword(userInput.Password) {
		log.Print("Password doesn't match")
		c.JSON(http.StatusUnauthorized, GetResponseDTO(401, nil, ErrorDTO{constants.PasswordNotMatch}))
		return
	}

	token := account.GenerateAccessToken()

	c.JSON(http.StatusOK, GetResponseDTO(200, LoginResponse{token}, ErrorDTO{}))

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// validate

}
