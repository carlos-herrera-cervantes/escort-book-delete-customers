package repositories

import (
	"context"

	"escort-book-delete-customers/models"
)

//go:generate mockgen -destination=./mocks/iescort_profile_repository.go -package=mocks --build_flags=--mod=mod . IEscortProfileRepository
type IEscortProfileRepository interface {
	Get(ctx context.Context, query string) (models.EscortProfile, error)
	Delete(ctx context.Context, query string) error
}
