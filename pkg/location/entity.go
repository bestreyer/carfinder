package location

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude" validate:"required"`
	Accuracy  float64 `json:"accuracy" validate:"required"`
	DriverId  int     `validate:"required"`
}