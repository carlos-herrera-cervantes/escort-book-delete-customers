package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ICardRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
