package repositories

import (
	"context"
	"escort-book-delete-customers/db"
	"escort-book-delete-customers/models"
)

type CustomerProfileRepository struct {
	Data *db.PostgresClient
}

func (r CustomerProfileRepository) Get(ctx context.Context, query string) (models.CustomerProfile, error) {
	row := r.Data.CustomerProfileDB.QueryRowContext(ctx, query)

	var customerProfile models.CustomerProfile

	if err := row.Scan(&customerProfile.CustomerId); err != nil {
		return customerProfile, err
	}

	return customerProfile, nil
}

func (r CustomerProfileRepository) Delete(ctx context.Context, query string) error {
	if _, err := r.Data.CustomerProfileDB.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
