package alerts

import (
	"fmt"
	"sync"
)

// AlertSystem handles the generation and dispatching of collision risk alerts
type AlertSystem struct {
	mu      sync.Mutex
	alerts  []string
	listeners []chan string
}

// NewAlertSystem creates and initializes a new AlertSystem instance
func NewAlertSystem() *AlertSystem {
	return &AlertSystem{
		alerts:    make([]string, 0),
		listeners: make([]chan string, 0),
	}
}

// SendAlert generates and dispatches an alert message
func (as *AlertSystem) SendAlert(alert string) error {
	as.mu.Lock()
	defer as.mu.Unlock()

	// Store the alert
	as.alerts = append(as.alerts, alert)

	// Dispatch to all listeners
	for _, listener := range as.listeners {
		select {
		case listener <- alert:
		default:
			// Non-blocking send - drop if channel is full
		}
	}

	return nil
}

// Subscribe adds a new listener channel for alert notifications
func (as *AlertSystem) Subscribe() chan string {
	as.mu.Lock()
	defer as.mu.Unlock()

	ch := make(chan string, 10) // Buffered channel
	as.listeners = append(as.listeners, ch)
	return ch
}

// GetAlerts returns a copy of all generated alerts
func (as *AlertSystem) GetAlerts() []string {
	as.mu.Lock()
	defer as.mu.Unlock()
	
	alerts := make([]string, len(as.alerts))
	copy(alerts, as.alerts)
	return alerts
}

// ClearAlerts removes all stored alerts
func (as *AlertSystem) ClearAlerts() {
	as.mu.Lock()
	defer as.mu.Unlock()
	as.alerts = make([]string, 0)
}