package handlers

import "github.com/gin-gonic/gin"

func ResponseSuccessHandler(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"code":   code,
		"status": "SUCCESS",
		"data":   data,
	})
}

func ResponseErrorHandler(c *gin.Context, code int, err interface{}) {
	c.JSON(code, gin.H{
		"code":   code,
		"status": "ERROR",
		"error":  err,
	})
}
