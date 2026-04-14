package config

import (
	"encoding/json"
	"os"
)

// Config holds the application configuration
type Config struct {
	Thresholds Thresholds `json:"thresholds"`
	Database   Database   `json:"database"`
}

// Thresholds contains various safety thresholds for monitoring
type Thresholds struct {
	WhaleDistanceAlert float64 `json:"whale_distance_alert"`
	WhaleDistanceDanger float64 `json:"whale_distance_danger"`
	VesselSpeedLimit   float64 `json:"vessel_speed_limit"`
	AlertWindowMinutes int     `json:"alert_window_minutes"`
}

// Database holds database connection configuration
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Thresholds.WhaleDistanceAlert <= 0 {
		return &ValidationError{"whale_distance_alert must be positive"}
	}
	if c.Thresholds.WhaleDistanceDanger <= 0 {
		return &ValidationError{"whale_distance_danger must be positive"}
	}
	if c.Thresholds.VesselSpeedLimit < 0 {
		return &ValidationError{"vessel_speed_limit cannot be negative"}
	}
	if c.Thresholds.AlertWindowMinutes <= 0 {
		return &ValidationError{"alert_window_minutes must be positive"}
	}
	if c.Database.Host == "" {
		return &ValidationError{"database host cannot be empty"}
	}
	if c.Database.Port <= 0 {
		return &ValidationError{"database port must be positive"}
	}
	if c.Database.Name == "" {
		return &ValidationError{"database name cannot be empty"}
	}
	if c.Database.User == "" {
		return &ValidationError{"database user cannot be empty"}
	}
	return nil
}

// ValidationError represents a configuration validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}