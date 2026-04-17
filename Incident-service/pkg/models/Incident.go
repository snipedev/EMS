package models

import (
	"time"
)

type Incident struct {
	ID               uint       `json:"id" Gorm:"primaryKey"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Status           string     `json:"status" Gorm:"default:open"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	ClosedAt         *time.Time `json:"closed_at"`
	ClosureStatement string     `json:"closure_statement" Gorm:"-"`
}

func (s *Incident) CloseEvent() {
	s.Status = "closed"
	s.ClosedAt = new(time.Now())
	s.ClosureStatement = "The event was closed successfully"
	// todo also add an event in the outbox table
}

func CreateIncident(title string, description string, status string) *Incident {
	return &Incident{
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   new(time.Now()),
		UpdatedAt:   nil,
		ClosedAt:    nil,
	}
}
