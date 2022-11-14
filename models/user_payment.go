package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPayment struct {
	Id              primitive.ObjectID `bson:"_id"`
	UserId          primitive.ObjectID `bson:"userId"`
	PaymentMethodId primitive.ObjectID `bson:"paymentMethodId"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}
