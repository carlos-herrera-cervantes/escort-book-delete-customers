package repositories

import (
	"context"
	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"

	"go.mongodb.org/mongo-driver/bson"
)

type UserPaymentRepository struct {
	Data *db.MongoClient
}

var userPaymentCollection = config.InitializeMongo().Collections.PaymentUserPayment

func (r UserPaymentRepository) DeleteMany(ctx context.Context, filter bson.M) error {
	collection := r.Data.PamentDB.Collection(userPaymentCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
