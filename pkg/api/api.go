package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandleStatus handles the status endpoint
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "running"}`))
}

// SetupRoutes registers all API routes
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/status", HandleStatus).Methods("GET")
}