package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/iaccess_token_repository.go -package=mocks --build_flags=--mod=mod . IAccessTokenRepository
type IAccessTokenRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
