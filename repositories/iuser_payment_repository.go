package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/iuser_payment_repository.go -package=mocks --build_flags=--mod=mod . IUserPaymentRepository
type IUserPaymentRepository interface {
	DeleteMany(ctx context.Context, filter bson.M) error
}
