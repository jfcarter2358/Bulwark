package api

import (
	"bulwark/constants"
	"bulwark/health"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alive(c *gin.Context) {
	if health.IsAlive {
		c.JSON(http.StatusOK, gin.H{"version": constants.VERSION})
		return
	}

	c.Status(http.StatusServiceUnavailable)
}

func Healthy(c *gin.Context) {
	if health.IsHealthy {
		c.JSON(http.StatusOK, gin.H{"version": constants.VERSION})
		return
	}

	c.Status(http.StatusServiceUnavailable)
}

func Ready(c *gin.Context) {
	if health.IsReady {
		c.JSON(http.StatusOK, gin.H{"version": constants.VERSION})
		return
	}

	c.Status(http.StatusServiceUnavailable)
}
