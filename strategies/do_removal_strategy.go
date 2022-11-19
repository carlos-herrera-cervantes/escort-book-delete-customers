package strategies

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"
	"escort-book-delete-customers/types"

	log "github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoRemovalStrategy struct {
	CustomerRemovalRepository repositories.ICustomerRemovalRepository
	CustomerProfileRepository repositories.ICustomerProfileRepository
	EscortProfileRepository   repositories.IEscortProfileRepository
}

var logger = log.New("strategies")

func (s DoRemovalStrategy) SwitchAccountRemoval(ctx context.Context, value []byte) {
	var userDeleteAccountEvent types.UserDeleteAccountEvent

	if err := json.Unmarshal(value, &userDeleteAccountEvent); err != nil {
		return
	}

	userId, _ := primitive.ObjectIDFromHex(userDeleteAccountEvent.UserId)
	removal := models.CustomerRemoval{
		UserId:      userId,
		UserEmail:   userDeleteAccountEvent.UserEmail,
		UserType:    userDeleteAccountEvent.UserType,
		ScheduledAt: time.Now().UTC().Add(672 * time.Hour),
	}
	removal.Validate()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		var err error

		if userDeleteAccountEvent.UserType == "Customer" {
			err = s.runForCustomer(ctx, userDeleteAccountEvent.UserId)
		} else {
			err = s.runForEscort(ctx, userDeleteAccountEvent.UserId)
		}

		if err != nil {
			logger.Warn(fmt.Sprintf("User: %s does not exist. Error: %s", userDeleteAccountEvent.UserEmail, err.Error()))
			return
		}

		if _, err := s.CustomerRemovalRepository.Create(ctx, removal); err != nil {
			logger.Error("Error saving a removal: ", err.Error())
		}
	}()

	wg.Wait()
}

func (s DoRemovalStrategy) runForCustomer(ctx context.Context, userId string) error {
	query := fmt.Sprintf("SELECT customer_id FROM profile WHERE customer_id = '%s'", userId)

	if _, err := s.CustomerProfileRepository.Get(ctx, query); err != nil {
		return err
	}

	return nil
}

func (s DoRemovalStrategy) runForEscort(ctx context.Context, userId string) error {
	query := fmt.Sprintf("SELECT escort_id FROM profile WHERE escort_id = '%s'", userId)

	if _, err := s.EscortProfileRepository.Get(ctx, query); err != nil {
		return err
	}

	return nil
}
