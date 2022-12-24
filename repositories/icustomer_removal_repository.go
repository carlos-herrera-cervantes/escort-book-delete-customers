package repositories

import (
	"context"

	"escort-book-delete-customers/models"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/icustomer_removal_repository.go -package=mocks --build_flags=--mod=mod . ICustomerRemovalRepository
type ICustomerRemovalRepository interface {
	Get(ctx context.Context, filter bson.M) (models.CustomerRemoval, error)
	GetMany(ctx context.Context, filter bson.M, offset, limit int64) ([]models.CustomerRemoval, error)
	Create(ctx context.Context, removal models.CustomerRemoval) (models.CustomerRemoval, error)
	Update(ctx context.Context, filter bson.M, document interface{}) (models.CustomerRemoval, error)
	Delete(ctx context.Context, filter bson.M) error
}
