package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	DB := db.Connect()
	DB.AutoMigrate(&models.Incident{})
	service := internals.NewIncidentService(DB)

	val := os.Getenv("API_KEY")

	log.Println("API_KEY: ", val)

	handler := handlers.NewIncidentHandler(service)

	// added for v1 endpoints
	AddRoutes(r, handler)

	r.Run(":8080")

}
