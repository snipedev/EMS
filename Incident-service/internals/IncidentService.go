package internals

import (
	"errors"
	"strings"

	"github.com/snipedev/ems/Incident-service/pkg/models"
	"gorm.io/gorm"
)

type IncidentService struct {
	DB *gorm.DB
}

func NewIncidentService(db *gorm.DB) *IncidentService {
	return &IncidentService{db}
}

func (s *IncidentService) CreateIncident(incident models.Incident) (models.Incident, error) {
	s.DB.Create(&incident)
	return incident, nil
}

func (s *IncidentService) GetIncident(id int) (*models.Incident, error) {
	var incident models.Incident
	err := s.DB.First(&incident, id).Error
	if err != nil {
		return nil, err
	}
	return &incident, nil
}

func (s *IncidentService) GetIncidents() ([]*models.Incident, error) {
	var incidents []*models.Incident

	err := s.DB.Find(&incidents).Error
	if err != nil {
		return nil, err
	}
	return incidents, nil
}

func (s *IncidentService) CloseIncident(id int) (*models.Incident, error) {
	incident, err := s.GetIncident(id)

	if err != nil {
		return nil, err
	}

	if incident != nil && strings.ToLower(incident.Status) == "closed" {
		return nil, errors.New("INCIDENT ALREADY CLOSED")
	}

	incident.CloseEvent()
	s.DB.Save(incident)
	return incident, nil
}
