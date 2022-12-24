package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUserRepositoryDeleteMany(t *testing.T) {
	userRepository := UserRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := userRepository.DeleteMany(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := userRepository.DeleteMany(ctxWithCancel, bson.M{})

		assert.Error(t, err)
	})
}
