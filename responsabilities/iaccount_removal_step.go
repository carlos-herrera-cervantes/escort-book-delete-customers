package responsabilities

import (
	"context"

	"escort-book-delete-customers/models"
)

type IAccountRemovalStep interface {
	Execute(ctx context.Context, removal []models.CustomerRemoval)
}
