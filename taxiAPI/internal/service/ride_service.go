package service

import (
	"awesomeProject/taxiAPI/internal/entity"
	customErrors "awesomeProject/taxiAPI/internal/errors"
	"github.com/google/uuid"
)

type RideStore interface {
	SaveRide(ride *entity.Ride) error
	FindRideByID(id string) (*entity.Ride, error)
	UpdateRideStatus(rideID string, status entity.Status) error
	AssignDriverToRide(rideID string, driverID string) error
	DeleteRide(id string) error
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
	id := uuid.New().String()
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
		ID:          id,
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

}

func (s *RideService) AssignDriverToRide(rideID string, driverID string) error {

}

func (s *RideService) DeleteRide(rideID string) error {

}
