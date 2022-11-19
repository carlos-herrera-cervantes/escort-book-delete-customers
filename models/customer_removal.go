package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerRemoval struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id"`
	UserType    string             `bson:"user_type"`
	UserEmail   string             `bson:"user_email"`
	Executed    bool               `bson:"executed"`
	ScheduledAt time.Time          `bson:"scheduled_date"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (cr *CustomerRemoval) Validate() error {
	cr.CreatedAt = time.Now().UTC()
	cr.UpdatedAt = time.Now().UTC()
	return nil
}
