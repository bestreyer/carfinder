package location

import "time"

type Location struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	DriverId  int
	UpdatedAt time.Time
}

type UpdateLocation struct {
	Location
	Accuracy float64 `json:"accuracy" validate:"required"`
}

type LocationWithDistance struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Id        int     `json:"id"`
	Distance  float64 `json:"distance"`
}

type DriverAroundLocation struct {
	Latitude  float64 `form:"latitude" validate:"required"`
	Longitude float64 `form:"longitude" validate:"required"`
	Radius    float64 `form:"radius"`
	Limit     int     `form:"limit"`
}
