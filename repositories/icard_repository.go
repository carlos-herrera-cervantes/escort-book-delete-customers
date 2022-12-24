package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/icard_repository.go -package=mocks --build_flags=--mod=mod . ICardRepository
type ICardRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
