package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUserPaymentRepositoryDeleteMany(t *testing.T) {
	userPaymentRepository := UserPaymentRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := userPaymentRepository.DeleteMany(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := userPaymentRepository.DeleteMany(ctxWithCancel, bson.M{})

		assert.Error(t, err)
	})
}
