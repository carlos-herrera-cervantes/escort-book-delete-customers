package models

import "time"

type CustomerProfile struct {
	Id         string
	CustomerId string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
