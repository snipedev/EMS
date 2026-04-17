package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snipedev/ems/Incident-service/handlers"
)

func AddRoutes(router *gin.Engine, handler *handlers.IncidentHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/incidents", handler.LogIncident)
		v1.GET("/incidents", handler.ListIncidents)
		v1.GET("/incidents/:id", handler.GetIncident)
		v1.PATCH("/incidents/:id", handler.CloseIncident)
	}
}
