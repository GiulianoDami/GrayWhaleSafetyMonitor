package models

import (
	"time"
)

// WhalePosition represents the location and movement data of a gray whale
type WhalePosition struct {
	ID        string    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Speed     float64   `json:"speed"`
	Heading   float64   `json:"heading"`
	Timestamp time.Time `json:"timestamp"`
}

// Vessel represents a maritime vessel with its current status and metadata
type Vessel struct {
	MMSI           string    `json:"mmsi"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Length         float64   `json:"length"`
	Speed          float64   `json:"speed"`
	Heading        float64   `json:"heading"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	LastUpdate     time.Time `json:"last_update"`
	EstimatedArrival time.Time `json:"estimated_arrival,omitempty"`
}

// RiskAssessment represents the calculated risk level for a vessel-whale interaction
type RiskAssessment struct {
	VesselID        string    `json:"vessel_id"`
	WhaleID         string    `json:"whale_id"`
	RiskLevel       string    `json:"risk_level"` // "low", "medium", "high", "critical"
	ConfidenceScore float64   `json:"confidence_score"`
	Distance        float64   `json:"distance"`
	ETA             time.Time `json:"eta"`
	Timestamp       time.Time `json:"timestamp"`
	AlertSent       bool      `json:"alert_sent"`
}