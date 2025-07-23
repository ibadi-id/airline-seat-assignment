package domain

import "time"

type Voucher struct {
	ID           int
	CrewName     string
	CrewID       string
	FlightNumber string
	FlightDate   string
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
