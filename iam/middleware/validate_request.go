package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/megastore/constants"
	"github.com/tanhaok/megastore/db"
	"github.com/tanhaok/megastore/handlers"
	"github.com/tanhaok/megastore/utils"
	"net/http"
)

func ValidateRequest(c *gin.Context) {
	// get api token from header
	apiToken := c.GetHeader(constants.ApiTokenRequestHeader)
	userId := c.GetHeader(constants.ApiUserIdRequestHeader)

	if apiToken == "" || userId == "" {
		handlers.ResponseErrorHandler(c, http.StatusUnauthorized, constants.MissingCredentials, nil)
		return
	}

	// get bearer token from redis
	hashMD5 := utils.ComputeMD5([]string{userId})
	accessToken, err := db.GetDataFromKey(fmt.Sprintf("%s_%s", hashMD5, apiToken))
	if err != nil || accessToken == nil || accessToken == "" {
		handlers.ResponseErrorHandler(c, http.StatusUnauthorized, constants.TokenNotFount, accessToken)
		return
	}

	isValidToken, userId, username, role := utils.ExtractDataFromToken(accessToken.(string))
	if !isValidToken {
		handlers.ResponseErrorHandler(c, http.StatusUnauthorized, constants.TokenNotFount, apiToken)
		return
	}

	c.Header(constants.ApiUserIdRequestHeader, userId)
	c.Header(constants.ApiUserRoles, role)
	c.Header(constants.ApiUserRequestHeader, username)

	c.Set(constants.ApiUserRoles, role)
	c.Next()
}
