package jobs

import (
	"context"
	"fmt"
	"sync"
	"time"

	"escort-book-delete-customers/repositories"
	"escort-book-delete-customers/responsabilities"

	log "github.com/inconshreveable/log15"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountRemovalJob struct {
	CustomerRemovalRepository repositories.ICustomerRemovalRepository
	CustomerProfileRepository repositories.ICustomerProfileRepository
	EscortProfileRepository   repositories.IEscortProfileRepository
	CardRepository            repositories.ICardRepository
	AccessTokenRepository     repositories.IAccessTokenRepository
	BankAccountRepository     repositories.IBankAccountRepository
	UserPaymentRepository     repositories.IUserPaymentRepository
	UserRepository            repositories.IUserRepository
}

type stopper func()

var logger = log.New("jobs")

func (j AccountRemovalJob) StartRemovalAccount(scheduleExpression string) stopper {
	c := cron.New()
	stopFn := func() { c.Stop() }

	c.AddFunc(scheduleExpression, func() {
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			j.removeCustomers(ctxWithTimeout)
		}()

		go func() {
			defer wg.Done()
			j.removeEscorts(ctxWithTimeout)
		}()

		wg.Wait()
	})
	c.Start()

	return stopFn
}

func (j AccountRemovalJob) removeEscorts(ctx context.Context) {
	now := time.Now().UTC().Add(24 * time.Hour)
	removalFilter := bson.M{
		"scheduled_date": bson.M{"$lte": primitive.NewDateTimeFromTime(now)},
		"user_type":      "Escort",
		"executed":       false,
	}

	removals, err := j.CustomerRemovalRepository.GetMany(ctx, removalFilter, int64(0), int64(1000))

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting removals: %s", err.Error()))
		return
	}

	if len(removals) == 0 {
		logger.Info("No deletions for escorts")
		return
	}

	customerRemovalStep := responsabilities.CustomerRemovalStep{
		CustomerRemovalRepository: j.CustomerRemovalRepository,
	}
	accessTokenStep := responsabilities.AccessTokenStep{
		AccessTokenRepository: j.AccessTokenRepository,
		Next:                  customerRemovalStep,
	}
	userStep := responsabilities.UserStep{
		UserRepository: j.UserRepository,
		Next:           accessTokenStep,
	}
	userPaymentStep := responsabilities.UserPaymentStep{
		UserPaymentRepository: j.UserPaymentRepository,
		Next:                  userStep,
	}
	bankAccountStep := responsabilities.EscortBankAccountStep{
		BankAccountRepository: j.BankAccountRepository,
		Next:                  userPaymentStep,
	}
	escortProfileStep := responsabilities.EscortProfileStep{
		EscortProfileRepository: j.EscortProfileRepository,
		Next:                    bankAccountStep,
	}
	escortProfileStep.Execute(ctx, removals)
}

func (j AccountRemovalJob) removeCustomers(ctx context.Context) {
	now := time.Now().UTC().Add(24 * time.Hour)
	removalFilter := bson.M{
		"scheduled_date": bson.M{"$lte": primitive.NewDateTimeFromTime(now)},
		"user_type":      "Customer",
		"executed":       false,
	}
	removals, err := j.CustomerRemovalRepository.GetMany(ctx, removalFilter, int64(0), int64(1000))

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting removals: %s", err.Error()))
		return
	}

	if len(removals) == 0 {
		logger.Info("No deletions for customers")
		return
	}

	customerRemovalStep := responsabilities.CustomerRemovalStep{
		CustomerRemovalRepository: j.CustomerRemovalRepository,
	}
	accessTokenStep := responsabilities.AccessTokenStep{
		AccessTokenRepository: j.AccessTokenRepository,
		Next:                  customerRemovalStep,
	}
	userStep := responsabilities.UserStep{
		UserRepository: j.UserRepository,
		Next:           accessTokenStep,
	}
	userPaymentStep := responsabilities.UserPaymentStep{
		UserPaymentRepository: j.UserPaymentRepository,
		Next:                  userStep,
	}
	bankAccountStep := responsabilities.CustomerBankAccountStep{
		BankAccountRepository: j.BankAccountRepository,
		Next:                  userPaymentStep,
	}
	cardStep := responsabilities.CardStep{
		CardRepository: j.CardRepository,
		Next:           bankAccountStep,
	}
	customerProfileStep := responsabilities.CustomerProfileStep{
		CustomerProfileRepository: j.CustomerProfileRepository,
		Next:                      cardStep,
	}
	customerProfileStep.Execute(ctx, removals)
}
