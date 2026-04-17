package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/snipedev/ems/Incident-service/handlers"
	"github.com/snipedev/ems/Incident-service/internals"
	"github.com/snipedev/ems/Incident-service/pkg/db"
	"github.com/snipedev/ems/Incident-service/pkg/models"
)

func loadconfig() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func AddRoutes(router *gin.Engine, handler *handlers.IncidentHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/incidents", handler.LogIncident)
		v1.GET("/incidents", handler.ListIncidents)
		v1.GET("/incidents/:id", handler.GetIncident)
	}
}

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	db := db.Connect()
	db.AutoMigrate(&models.Incident{})
	service := internals.NewIncidentService(db)

	val := os.Getenv("API_KEY")

	log.Println("API_KEY: ", val)

	handler := handlers.NewIncidentHandler(service)
	AddRoutes(r, handler)

	// added for v1 endpoints

	r.Run(":8080")

}
