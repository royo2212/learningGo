package entity

type Driver struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	IsAvailable  bool   `json:"is_available"`
	CarType      string `json:"car_type"`
	LicensePlate string `json:"license_plate"`
}
