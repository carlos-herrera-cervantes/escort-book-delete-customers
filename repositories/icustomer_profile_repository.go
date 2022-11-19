package repositories

import (
	"context"
	"escort-book-delete-customers/models"
)

type ICustomerProfileRepository interface {
	Get(ctx context.Context, query string) (models.CustomerProfile, error)
	Delete(ctx context.Context, query string) error
}
