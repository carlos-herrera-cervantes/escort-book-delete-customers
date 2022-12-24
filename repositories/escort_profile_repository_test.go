package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"

	"github.com/stretchr/testify/assert"
)

func TestEscortProfileRepositoryGet(t *testing.T) {
	escortProfileRepository := EscortProfileRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		query := "SELECT * FROM profile WHERE id = 'dummy id'"
		_, err := escortProfileRepository.Get(context.Background(), query)
		assert.Error(t, err)
	})
}

func TestEscortProfileRepositoryDelete(t *testing.T) {
	escortProfileRepository := EscortProfileRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		query := "DELETE FROM profile WHERE id = 'dummy id'"
		err := escortProfileRepository.Delete(context.Background(), query)
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		query := "DELETE FROM profile WHERE id = 'dummy id'"
		err := escortProfileRepository.Delete(ctxWithCancel, query)

		assert.Error(t, err)
	})
}