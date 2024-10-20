package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/megastore/logging"
	"go.uber.org/zap"
)

func ResponseSuccessHandler(c *gin.Context, code int, data interface{}) {
	logging.LOGGER.Error("Success to handle request",
		zap.String("endpoint", c.Request.RequestURI),
		zap.String("method", c.Request.Method),
		zap.String("remote_address", c.Request.RemoteAddr),
		zap.Any("header", c.Request.Header),
	)

	c.JSON(code, gin.H{
		"code":   code,
		"status": "SUCCESS",
		"data":   data,
	})
}

func ResponseErrorHandler(c *gin.Context, code int, err interface{}, traceData any) {
	logging.LOGGER.Error("Failed to handle request",
		zap.String("endpoint", c.Request.RequestURI),
		zap.String("method", c.Request.Method),
		zap.String("remote_address", c.Request.RemoteAddr),
		zap.Any("data_in_request", traceData),
		zap.Any("header", c.Request.Header),
		zap.Any("error", err),
	)
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"status": "ERROR",
		"error":  err,
	})
	c.Abort()
}
