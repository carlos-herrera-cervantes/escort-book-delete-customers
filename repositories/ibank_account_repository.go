package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type IBankAccountRepository interface {
	DeleteCustomerBankAccounts(ctx context.Context, filter bson.M) error
	DeleteEscortBankAccounts(ctx context.Context, filter bson.M) error
}
