package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"
	"escort-book-delete-customers/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCustomerRemovalRepositoryGet(t *testing.T) {
	customerRemovalRepository := CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := customerRemovalRepository.Get(context.Background(), bson.M{})
		assert.Error(t, err)
	})
}

func TestCustomerRemovalRepositoryGetMany(t *testing.T) {
	customerRemovalRepository := CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return empty slice", func(t *testing.T) {
		deletions, err := customerRemovalRepository.GetMany(
			context.Background(), bson.M{}, 0, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, deletions)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := customerRemovalRepository.GetMany(
			ctxWithCancel, bson.M{}, 0, 10,
		)

		assert.Error(t, err)
	})
}

func TestCustomerRemovalRepositoryCreate(t *testing.T) {
	customerRemovalRepository := CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := customerRemovalRepository.Create(ctxWithCancel, models.CustomerRemoval{})

		assert.Error(t, err)
	})
}

func TestCustomerRemovalRepositoryUpdate(t *testing.T) {
	customerRemovalRepository := CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := customerRemovalRepository.Update(
			context.Background(),
			bson.M{"user_email": "test@example.com"},
			map[string]bool{
				"executed": true,
			},
		)
		assert.Error(t, err)
	})
}

func TestCustomerRemovalRepositoryDelete(t *testing.T) {
	customerRemovalRepository := CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := customerRemovalRepository.Delete(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := customerRemovalRepository.Delete(ctxWithCancel, bson.M{})

		assert.Error(t, err)
	})
}
