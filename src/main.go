// main.go

package main

import (
	"bulwark/buffer"
	"bulwark/config"
	"bulwark/queue"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	logger "github.com/jfcarter2358/go-logger"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config.LoadConfig()
	logger.SetLevel(config.Config.LogLevel)
	logger.SetFormat(config.Config.LogFormat)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.Config.TLSSkipVerify}

	router = gin.New()
	router.Use(gin.LoggerWithFormatter(logger.ConsoleLogFormatter))
	router.Use(gin.Recovery())

	logger.Infof("", "Running with port: %d", config.Config.Port)

	initializeRoutes()

	buffer.Init()
	queue.Init()

	rand.Seed(time.Now().UnixNano())

	routerPort := fmt.Sprintf(":%d", config.Config.Port)
	if config.Config.TLSEnabled {
		logger.Infof("", "Running with TLS loaded from %s and %s", config.Config.TLSCrtPath, config.Config.TLSKeyPath)
		router.RunTLS(routerPort, config.Config.TLSCrtPath, config.Config.TLSKeyPath)
	} else {
		router.Run(routerPort)
	}
}
