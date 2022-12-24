package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/iuser_repository.go -package=mocks --build_flags=--mod=mod . IUserRepository
type IUserRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
