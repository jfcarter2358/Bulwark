// routes.go

package main

import (
	"bulwark/api"
	"bulwark/middleware"
)

func initializeRoutes() {
	healthRoutes := router.Group("/health", middleware.CORSMiddleware())
	{
		healthRoutes.GET("/healthy", api.Healthy)
		healthRoutes.GET("/ready", api.Ready)
		healthRoutes.GET("/alive", api.Alive)
	}

	apiRoutes := router.Group("/api", middleware.CORSMiddleware())
	{
		v1Routes := apiRoutes.Group("/v1")
		{
			queueRoutes := v1Routes.Group("/queue")
			{
				queueRoutes.DELETE("/:name", middleware.EnsureAuth(), api.QueueDelete)
				queueRoutes.GET("/:name", middleware.EnsureAuth(), api.QueuePop)
				queueRoutes.POST("/:name", middleware.EnsureAuth(), api.QueueCreate)
				queueRoutes.PUT("/:name", middleware.EnsureAuth(), api.QueuePush)
				queueRoutes.PUT("/:name/ping", middleware.EnsureAuth(), api.QueuePing)
			}
			bufferRoutes := v1Routes.Group("/buffer")
			{
				bufferRoutes.DELETE("/:name", middleware.EnsureAuth(), api.BufferDelete)
				bufferRoutes.GET("/:name", middleware.EnsureAuth(), api.BufferGet)
				bufferRoutes.POST("/:name", middleware.EnsureAuth(), api.BufferCreate)
				bufferRoutes.PUT("/:name", middleware.EnsureAuth(), api.BufferSet)
				bufferRoutes.PUT("/:name/ping", middleware.EnsureAuth(), api.BufferPing)
			}
		}
	}
}
