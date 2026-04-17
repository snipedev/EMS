package internals

import (
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
