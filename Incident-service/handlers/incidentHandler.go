package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snipedev/ems/Incident-service/internals"
	"github.com/snipedev/ems/Incident-service/pkg/models"
)

type IncidentHandler struct {
	service *internals.IncidentService
}

func NewIncidentHandler(svc *internals.IncidentService) *IncidentHandler {
	return &IncidentHandler{service: svc}
}

func (handler *IncidentHandler) LogIncident(c *gin.Context) {
	var incidentRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	err := c.ShouldBindJSON(&incidentRequest)

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	incident := models.CreateIncident(
		incidentRequest.Title,
		incidentRequest.Description,
		"pending",
	)

	createdIncident, err := handler.service.CreateIncident(*incident)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSONP(http.StatusCreated, gin.H{
		"status":   "Created",
		"incident": createdIncident,
	})
}

func (handler *IncidentHandler) GetIncident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	incident, err := handler.service.GetIncident(id)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSONP(http.StatusOK, gin.H{
		"status":   "Incident",
		"incident": &incident,
	})
}

func (handler *IncidentHandler) ListIncidents(c *gin.Context) {
	//var incidents []*models.Incident
	incidents, err := handler.service.GetIncidents()

	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSONP(http.StatusOK, gin.H{
		"status":    "Incidents",
		"incidents": incidents,
	})
}
