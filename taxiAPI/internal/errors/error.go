package errors

import "errors"

var (
	ErrRideNotFound                       = errors.New("ride not found")
	ErrRideIDRequired                     = errors.New("ride ID is required")
	ErrOriginRequired                     = errors.New("origin is required")
	ErrDestinationRequired                = errors.New("destination is required")
	ErrPassengerIDRequired                = errors.New("passenger ID is required")
	ErrInvalidRideStatus                  = errors.New("invalid ride status")
	ErrCannotChangeCompletedRide          = errors.New("cannot change completed ride")
	ErrDriverIDRequired                   = errors.New("drive ID is required")
	ErrCannotAssignDriverToNonPendingRide = errors.New("cannot assign driver to non pending ride")
	ErrRideAlreadyAssigned                = errors.New("ride already assigned")
	ErrDriverAlreadyAssignedToRide        = errors.New("this driver already assigned to this ride")
)
