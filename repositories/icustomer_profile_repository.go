package repositories

import (
	"context"

	"escort-book-delete-customers/models"
)

//go:generate mockgen -destination=./mocks/icustomer_profile_repository.go -package=mocks --build_flags=--mod=mod . ICustomerProfileRepository
type ICustomerProfileRepository interface {
	Get(ctx context.Context, query string) (models.CustomerProfile, error)
	Delete(ctx context.Context, query string) error
}
