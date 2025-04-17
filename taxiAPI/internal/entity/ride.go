package entity

type Ride struct {
	ID          string `json:"id"`
	PassengerID string `json:"passenger_id"`
	DriverID    string `json:"driver_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Status      Status `json:"status"`
}
type Status string

const (
	StatusPending   Status = "pending"
	StatusAccepted  Status = "accepted"
	StatusCompleted Status = "completed"
	StatusCancelled Status = "cancelled"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusAccepted, StatusCompleted, StatusCancelled:
		return true
	}
	return false
}
