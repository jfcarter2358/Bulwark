// utils.go

package utils

import (
	logger "github.com/jfcarter2358/go-logger"

	"github.com/gin-gonic/gin"
)

func Error(err error, c *gin.Context, statusCode int) {
	logger.Error("", err.Error())
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
