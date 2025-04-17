package service

import (
	"taxiAPI/internal/entity"
	customErrors "taxiAPI/internal/errors"
)

type RideStore interface {
	SaveRide(ride *entity.Ride) error
	FindRideByID(id string) (*entity.Ride, error)
	UpdateRideStatus(rideID string, status entity.Status) error
	AssignDriverToRide(rideID string, driverID string) error
	GetAllRides() ([]*entity.Ride, error)
}
type RideService struct {
	store RideStore
}

func NewRideService(store RideStore) *RideService {
	return &RideService{
		store: store,
	}
}
func (s *RideService) CreateRide(passengerID, origin, destination string) (*entity.Ride, error) {
	if passengerID == "" {
		return nil, customErrors.ErrPassengerIDRequired
	}
	if origin == "" {
		return nil, customErrors.ErrOriginRequired
	}
	if destination == "" {
		return nil, customErrors.ErrDestinationRequired
	}
	ride := &entity.Ride{
		PassengerID: passengerID,
		Origin:      origin,
		Destination: destination,
		Status:      entity.StatusPending,
	}
	err := s.store.SaveRide(ride)
	if err != nil {
		return nil, err
	}
	return ride, nil
}
func (s *RideService) GetRide(rideID string) (*entity.Ride, error) {
	if rideID == "" {
		return nil, customErrors.ErrRideIDRequired
	}
	ride, err := s.store.FindRideByID(rideID)
	if err != nil {
		return nil, err
	}
	return ride, nil
}

func (s *RideService) GetAllRides() ([]*entity.Ride, error) {
	return s.store.GetAllRides()
}

func (s *RideService) UpdateRideStatus(rideID string, status entity.Status) error {
	if rideID == "" {
		return customErrors.ErrRideIDRequired
	}
	ride, err := s.store.FindRideByID(rideID)
	if err != nil {
		return err
	}
	if ride.Status == entity.StatusCompleted && status != entity.StatusCompleted {
		return customErrors.ErrCannotChangeCompletedRide
	}
	if !status.IsValid() {
		return customErrors.ErrInvalidRideStatus
	}
	ride.Status = status
	return s.store.UpdateRideStatus(rideID, ride.Status)
}

func (s *RideService) AssignDriverToRide(rideID string, driverID string) error {
	if rideID == "" {
		return customErrors.ErrRideIDRequired
	}
	if driverID == "" {
		return customErrors.ErrDriverIDRequired
	}
	ride, err := s.store.FindRideByID(rideID)
	if err != nil {
		return err
	}
	if ride.DriverID != "" {
		if ride.DriverID == driverID {
			return customErrors.ErrDriverAlreadyAssignedToRide
		}
		return customErrors.ErrRideAlreadyAssigned
	}
	if ride.Status != entity.StatusPending {
		return customErrors.ErrCannotAssignDriverToNonPendingRide
	}
	ride.DriverID = driverID
	ride.Status = entity.StatusAccepted
	return s.store.AssignDriverToRide(rideID, driverID)
}
