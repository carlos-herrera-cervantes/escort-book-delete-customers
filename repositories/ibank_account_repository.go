package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/ibank_account_repository.go -package=mocks --build_flags=--mod=mod . IBankAccountRepository
type IBankAccountRepository interface {
	DeleteCustomerBankAccounts(ctx context.Context, filter bson.M) error
	DeleteEscortBankAccounts(ctx context.Context, filter bson.M) error
}
