package responsabilities

import (
	"context"

	"escort-book-delete-customers/models"
)

//go:generate mockgen -destination=./mocks/iaccount_removal_step.go -package=mocks --build_flags=--mod=mod . IAccountRemovalStep
type IAccountRemovalStep interface {
	Execute(ctx context.Context, removal []models.CustomerRemoval)
}
