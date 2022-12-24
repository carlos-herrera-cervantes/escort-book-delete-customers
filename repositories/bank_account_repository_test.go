package repositories

import (
	"context"
	"testing"

	"escort-book-delete-customers/db"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBankAccountRepositoryDeleteCustomerBankAccounts(t *testing.T) {
	bankAccountRepository := BankAccountRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := bankAccountRepository.DeleteCustomerBankAccounts(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := bankAccountRepository.DeleteCustomerBankAccounts(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})
}

func TestBankAccountRepositoryDeleteEscortBankAccounts(t *testing.T) {
	bankAccountRepository := BankAccountRepository{
		Data: db.NewMongoClient(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := bankAccountRepository.DeleteEscortBankAccounts(context.Background(), bson.M{})
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := bankAccountRepository.DeleteEscortBankAccounts(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})
}
