package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	customErrors "taxiAPI/internal/errors"
	"taxiAPI/internal/service"
)

type RideHandler struct {
	service *service.RideService
}

func NewRideHandler(service *service.RideService) *RideHandler {
	return &RideHandler{
		service: service,
	}
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

type createRideRequest struct {
	PassengerID string `json:"passenger_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

func (h *RideHandler) CreateRide(w http.ResponseWriter, r *http.Request) {
	var req createRideRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	ride, err := h.service.CreateRide(req.PassengerID, req.Origin, req.Destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *RideHandler) GetRide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rideID := vars["id"]
	if rideID == "" {
		http.Error(w, customErrors.ErrRideIDRequired.Error(), http.StatusBadRequest)
		return
	}
	ride, err := h.service.GetRide(rideID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *RideHandler) GetAllRides(w http.ResponseWriter, r *http.Request) {
	rides, err := h.service.GetAllRides()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(rides); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *RideHandler) AssignDriverToRide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rideID := vars["id"]
	if rideID == "" {
		http.Error(w, customErrors.ErrRideIDRequired.Error(), http.StatusBadRequest)
		return
	}
	var req struct {
		DriverID string `json:"driver_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.AssignDriverToRide(rideID, req.DriverID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{
		"message":   "Driver assigned successfully",
		"ride_id":   rideID,
		"driver_id": req.DriverID,
	})
	if err != nil {
		return
	}
}
