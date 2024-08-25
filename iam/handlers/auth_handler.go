package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/megastore/constants"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/dto"
	"github.com/tanhaok/megastore/kafka"
	"github.com/tanhaok/megastore/models"
	"github.com/tanhaok/megastore/utils"
	"log"
	"net/http"
)

// ========================= Functions =========================

// Register new account
func Register(c *gin.Context) {
	var userInput dto.RegisterRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, userInput)
		return
	}

	if ok, errors := utils.ValidateInput(userInput); !ok {
		ResponseErrorHandler(c, http.StatusBadRequest, errors, userInput)
		return
	}

	if models.ExistsByEmailOrUsername(userInput.Email, userInput.Username) {
		ResponseErrorHandler(c, http.StatusBadRequest, fmt.Sprintf(constants.AccountExists, userInput.Email, userInput.Username), userInput)
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
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, account)
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
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, userInput)
		return
	}

	var account models.Account
	var err error

	if account, err = models.GetAccountByEmailOrUsername(userInput.Email, userInput.Username); err != nil {
		ResponseErrorHandler(c, http.StatusNotFound, constants.AccountNotFound, userInput)
		return
	}

	if !account.ComparePassword(userInput.Password) {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.PasswordNotMatch, userInput)
		return
	}

	token := account.GenerateAccessToken()
	ResponseSuccessHandler(c, http.StatusOK, dto.LoginResponse{ApiToken: token})

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// get api token from header
	apiToken := c.GetHeader(constants.ApiTokenRequestHeader)
	userId := c.GetHeader(constants.ApiUserIdRequestHeader)

	if apiToken == "" || userId == "" {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.Unauthorized, apiToken)
		return
	}

	// get bearer token from redis
	hashMD5 := utils.ComputeMD5([]string{userId})
	accessToken, err := db.GetDataFromKey(fmt.Sprintf("%s_%s", hashMD5, apiToken))
	if err != nil || accessToken == nil || accessToken == "" {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.TokenNotFount, accessToken)
		return
	}

	isValidToken, userId, username, role := utils.ExtractDataFromToken(accessToken.(string))
	if !isValidToken {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.TokenNotFount, apiToken)
		return
	}
	ResponseSuccessHandler(c, http.StatusOK, nil)
	c.Header(constants.ApiUserIdRequestHeader, userId)
	c.Header(constants.ApiUserRoles, role)
	c.Header(constants.ApiUserRequestHeader, username)

}
