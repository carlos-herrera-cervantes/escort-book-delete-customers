package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAccessTokenRepositoryDeleteMany(t *testing.T) {
	accessTokenRepository := AccessTokenRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := accessTokenRepository.DeleteMany(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := accessTokenRepository.DeleteMany(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})
}