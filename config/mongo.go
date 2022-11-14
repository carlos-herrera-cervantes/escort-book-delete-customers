package config

import "os"

type mongo struct {
	Host        string
	Databases   mongoDatabases
	Collections mongoCollections
}

type mongoDatabases struct {
	Payment    string
	Authorizer string
	Scheduler  string
}

type mongoCollections struct {
	AuthorizerAccessToken      string
	AuthorizerUser             string
	SchedulerCustomerRemoval   string
	PaymentCustomerBankAccount string
	PaymentEscortBankAccount   string
	PaymentCard                string
	PaymentUserPayment         string
}

var singleMongo *mongo

func InitializeMongo() *mongo {
	if singleMongo != nil {
		return singleMongo
	}

	lock.Lock()
	defer lock.Unlock()

	singleMongo = &mongo{
		Host: os.Getenv("MONGO_HOST"),
		Databases: mongoDatabases{
			Payment:    os.Getenv("ESCORT_BOOK_PAYMENT_DB"),
			Authorizer: os.Getenv("ESCORT_BOOK_AUTHORIZER_DB"),
			Scheduler:  os.Getenv("ESCORT_BOOK_SCHEDULER_DB"),
		},
		Collections: mongoCollections{
			AuthorizerAccessToken:      "accesstokens",
			AuthorizerUser:             "users",
			SchedulerCustomerRemoval:   "customer-removal",
			PaymentCustomerBankAccount: "customer_bank_accounts",
			PaymentEscortBankAccount:   "escort_bank_accounts",
			PaymentCard:                "cards",
			PaymentUserPayment:         "user_payments",
		},
	}

	return singleMongo
}
