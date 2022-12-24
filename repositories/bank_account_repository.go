package repositories

import (
	"context"

	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"

	"go.mongodb.org/mongo-driver/bson"
)

type BankAccountRepository struct {
	Data *db.MongoClient
}

var customerBankAccountCollection = config.InitializeMongo().Collections.PaymentCustomerBankAccount
var escortBankAccountCollection = config.InitializeMongo().Collections.PaymentEscortBankAccount

func (r BankAccountRepository) DeleteCustomerBankAccounts(ctx context.Context, filter bson.M) error {
	collection := r.Data.PamentDB.Collection(customerBankAccountCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (r BankAccountRepository) DeleteEscortBankAccounts(ctx context.Context, filter bson.M) error {
	collection := r.Data.PamentDB.Collection(escortBankAccountCollection)

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
