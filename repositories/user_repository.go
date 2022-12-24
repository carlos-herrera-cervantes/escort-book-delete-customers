package repositories

import (
	"context"

	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	Data *db.MongoClient
}

var userCollection = config.InitializeMongo().Collections.AuthorizerUser

func (r UserRepository) DeleteMany(ctx context.Context, filter bson.M) error {
	collection := r.Data.AuthorizerDB.Collection(userCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
