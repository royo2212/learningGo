package storage

import (
	"awesomeProject/taxiAPI/internal/entity"
	customErrors "awesomeProject/taxiAPI/internal/errors"
	"sync"
)

type RideMemory struct {
	mutex sync.RWMutex
	rides map[string]*entity.Ride
}

func NewRideMemory() *RideMemory {
	return &RideMemory{
		rides: make(map[string]*entity.Ride),
	}
}

func (store *RideMemory) SaveRide(ride *entity.Ride) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.rides[ride.ID] = ride
	return nil
}
func (store *RideMemory) FindRideByID(id string) (*entity.Ride, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	ride, ok := store.rides[id]
	if !ok {
		return nil, customErrors.ErrRideNotFound
	}
	return ride, nil
}
func (store *RideMemory) UpdateRideStatus(rideID string, status entity.Status) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	ride, ok := store.rides[rideID]
	if !ok {
		return customErrors.ErrRideNotFound
	}
	ride.Status = status
	return nil
}
func (store *RideMemory) AssignDriverToRide(rideID string, driverID string) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	ride, ok := store.rides[rideID]
	if !ok {
		return customErrors.ErrRideNotFound
	}
	ride.DriverID = driverID
	return nil
}
func (store *RideMemory) DeleteRide(id string) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	_, ok := store.rides[id]
	if !ok {
		return customErrors.ErrRideNotFound
	}
	delete(store.rides, id)
	return nil
}
func (store *RideMemory) GetAllRides() ([]*entity.Ride, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	rides := make([]*entity.Ride, 0, len(store.rides))
	for _, ride := range store.rides {
		rides = append(rides, ride)
	}
	return rides, nil
}
