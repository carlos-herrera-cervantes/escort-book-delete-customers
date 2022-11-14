package models

import "time"

type EscortProfile struct {
	Id        string
	EscortId  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
