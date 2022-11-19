package repositories

import (
	"context"
	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"

	"go.mongodb.org/mongo-driver/bson"
)

type AccessTokenRepository struct {
	Data *db.MongoClient
}

var accessTokenCollection = config.InitializeMongo().Collections.AuthorizerAccessToken

func (r AccessTokenRepository) DeleteMany(ctx context.Context, filter bson.M) error {
	collection := r.Data.AuthorizerDB.Collection(accessTokenCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
