package repositories

import (
	"context"
	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"

	"go.mongodb.org/mongo-driver/bson"
)

type CardRepository struct {
	Data *db.MongoClient
}

var cardCollection = config.InitializeMongo().Collections.PaymentCard

func (r CardRepository) DeleteMany(ctx context.Context, filter bson.M) error {
	collection := r.Data.PamentDB.Collection(cardCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
