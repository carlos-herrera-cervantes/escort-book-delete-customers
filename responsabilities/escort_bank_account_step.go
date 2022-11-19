package responsabilities

import (
	"context"
	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type EscortBankAccountStep struct {
	BankAccountRepository repositories.IBankAccountRepository
	Next                  IAccountRemovalStep
}

func (s EscortBankAccountStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.BankAccountRepository.DeleteEscortBankAccounts(
		ctx,
		bson.M{"escortId": bson.M{"$in": listOfUserIds}},
	); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting escort bank accounts: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
