package middleware

import (
	"bulwark/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jfcarter2358/go-logger"
)

func EnsureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Bulwark-API")
		if token == "" {
			logger.Debugf("", "Token is empty")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token != config.Config.SecretKey {
			logger.Debugf("", "Got token %s", token)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
