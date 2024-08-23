package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/MyStore/constants"
	"github.com/tanhaok/MyStore/db"
	"github.com/tanhaok/MyStore/dto"
	"github.com/tanhaok/MyStore/kafka"
	"github.com/tanhaok/MyStore/models"
	"github.com/tanhaok/MyStore/utils"
	"log"
	"net/http"
)

// ========================= Functions =========================

// Register new account
func Register(c *gin.Context) {
	var userInput dto.RegisterRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson)
		return
	}

	if ok, errors := utils.ValidateInput(userInput); !ok {
		ResponseErrorHandler(c, http.StatusBadRequest, errors)
		return
	}

	if models.ExistsByEmailOrUsername(userInput.Email, userInput.Username) {
		ResponseErrorHandler(c, http.StatusBadRequest, fmt.Sprintf(constants.AccountExists, userInput.Email, userInput.Username))
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
		log.Printf("Error when saving account. %v", err)
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson)
		return
	}

	// send msg to kafka
	serializedMessage := account.GetSerializedMessageForActiveNewUser()
	_ = db.SaveActiveTokenToCache(account.Username, serializedMessage)

	kafka.PushMessageNewUser(serializedMessage)

	ResponseSuccessHandler(c, http.StatusCreated, nil)
}

// Login verify user credentials and return uuid pair with token saved in redis
func Login(c *gin.Context) {
	var userInput dto.LoginRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Printf("Value is incorrect. %v", err)
		c.JSON(http.StatusBadRequest, dto.GetResponseDTO(404, nil, dto.ErrorDTO{constants.MessageErrorBindJson}))
		return
	}

	var account models.Account
	var err error

	if account, err = models.GetAccountByEmailOrUsername(userInput.Email, userInput.Username); err != nil {
		log.Printf("Account donesn't exits")
		c.JSON(http.StatusNotFound, dto.GetResponseDTO(404, nil, dto.ErrorDTO{constants.AccountNotFound}))
		return
	}

	if !account.ComparePassword(userInput.Password) {
		log.Print("Password doesn't match")
		c.JSON(http.StatusUnauthorized, dto.GetResponseDTO(401, nil, dto.ErrorDTO{Message: constants.PasswordNotMatch}))
		return
	}

	token := account.GenerateAccessToken()

	c.JSON(http.StatusOK, dto.GetResponseDTO(200, dto.LoginResponse{ApiToken: token}, dto.ErrorDTO{}))

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// validate

}
