package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerBankAccountStep struct {
	BankAccountRepository repositories.IBankAccountRepository
	Next                  IAccountRemovalStep
}

func (s CustomerBankAccountStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.BankAccountRepository.DeleteCustomerBankAccounts(
		ctx,
		bson.M{"customerId": bson.M{"$in": listOfUserIds}},
	); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting customer bank accounts: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
