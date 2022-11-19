package repositories

import (
	"context"
	"escort-book-delete-customers/models"
)

type IEscortProfileRepository interface {
	Get(ctx context.Context, query string) (models.EscortProfile, error)
	Delete(ctx context.Context, query string) error
}
