package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type UserPaymentStep struct {
	UserPaymentRepository repositories.IUserPaymentRepository
	Next                  IAccountRemovalStep
}

func (s UserPaymentStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.UserPaymentRepository.DeleteMany(ctx, bson.M{"userId": bson.M{"$in": listOfUserIds}}); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting customer payments: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
