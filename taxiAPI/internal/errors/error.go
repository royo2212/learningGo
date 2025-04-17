package errors

import "errors"

var (
	ErrRideNotFound        = errors.New("ride not found")
	ErrRideIDRequired      = errors.New("ride ID is required")
	ErrOriginRequired      = errors.New("origin is required")
	ErrDestinationRequired = errors.New("destination is required")
	ErrPassengerIDRequired = errors.New("passenger ID is required")
)
