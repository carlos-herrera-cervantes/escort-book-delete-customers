package repositories

import (
	"context"

	"escort-book-delete-customers/db"
	"escort-book-delete-customers/models"
)

type EscortProfileRepository struct {
	Data *db.PostgresClient
}

func (r EscortProfileRepository) Get(ctx context.Context, query string) (models.EscortProfile, error) {
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var escortProfile models.EscortProfile

	if err := row.Scan(&escortProfile.EscortId); err != nil {
		return escortProfile, err
	}

	return escortProfile, nil
}

func (r EscortProfileRepository) Delete(ctx context.Context, query string) error {
	if _, err := r.Data.EscortProfileDB.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
