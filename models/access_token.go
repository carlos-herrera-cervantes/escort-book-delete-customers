package models

import "time"

type AccessToken struct {
	Id        string    `bson:"_id"`
	User      string    `bson:"user"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
