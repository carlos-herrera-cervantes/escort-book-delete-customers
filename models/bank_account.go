package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EscortBankAccount struct {
	Id        primitive.ObjectID `bson:"_id"`
	EscortId  primitive.ObjectID `bson:"escortId"`
	Clabe     string             `bson:"clabe"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type CustomerBankAccount struct {
	Id         primitive.ObjectID `bson:"_id"`
	CustomerId primitive.ObjectID `bson:"customerId"`
	Clabe      string             `bson:"clabe"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}
