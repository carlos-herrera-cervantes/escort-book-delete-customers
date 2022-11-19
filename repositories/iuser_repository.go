package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
