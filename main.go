package main

import (
	"escort-book-delete-customers/config"
	"escort-book-delete-customers/consumers"
	"escort-book-delete-customers/db"
	"escort-book-delete-customers/handlers"
	"escort-book-delete-customers/jobs"
	"escort-book-delete-customers/repositories"
	"escort-book-delete-customers/strategies"
)

func main() {
	customerRemovalRepository := &repositories.CustomerRemovalRepository{
		Data: db.NewMongoClient(),
	}
	customerProfileRepository := &repositories.CustomerProfileRepository{
		Data: db.NewPostgresClient(),
	}
	escortProfileRepository := &repositories.EscortProfileRepository{
		Data: db.NewPostgresClient(),
	}
	cardRepository := &repositories.CardRepository{
		Data: db.NewMongoClient(),
	}
	accessTokenRepository := &repositories.AccessTokenRepository{
		Data: db.NewMongoClient(),
	}
	bankAccountRepository := &repositories.BankAccountRepository{
		Data: db.NewMongoClient(),
	}
	userPaymentRepository := &repositories.UserPaymentRepository{
		Data: db.NewMongoClient(),
	}
	userRepository := &repositories.UserRepository{
		Data: db.NewMongoClient(),
	}

	doRemovalStrategy := &strategies.DoRemovalStrategy{
		CustomerRemovalRepository: customerRemovalRepository,
		CustomerProfileRepository: customerProfileRepository,
		EscortProfileRepository:   escortProfileRepository,
	}
	undoRemovalStrategy := &strategies.UndoRemovalStrategy{
		CustomerRemovalRepository: customerRemovalRepository,
	}
	accountStrategies := map[string]strategies.IAccountStrategy{
		config.InitializeKafka().Topics.UserDeleteAccount: doRemovalStrategy,
		config.InitializeKafka().Topics.UserActiveAccount: undoRemovalStrategy,
	}

	accountHandler := handlers.AccountHandler{
		StrategyManager: strategies.AccountStrategyManager{
			Strategies: accountStrategies,
		},
	}

	accountRemovalJob := jobs.AccountRemovalJob{
		CustomerRemovalRepository: customerRemovalRepository,
		CustomerProfileRepository: customerProfileRepository,
		EscortProfileRepository:   escortProfileRepository,
		CardRepository:            cardRepository,
		AccessTokenRepository:     accessTokenRepository,
		BankAccountRepository:     bankAccountRepository,
		UserPaymentRepository:     userPaymentRepository,
		UserRepository:            userRepository,
	}
	accountRemovalJobStopper := accountRemovalJob.StartRemovalAccount(
		config.InitializeJob().RemovalAccount,
	)

	accountRemovalConsumer := consumers.AccountRemovalConsumer{
		EventHandler: accountHandler,
	}
	accountRemovalConsumer.StartConsumer()

	accountRemovalJobStopper()
}
