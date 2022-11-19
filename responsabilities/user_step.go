package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type UserStep struct {
	UserRepository repositories.IUserRepository
	Next           IAccountRemovalStep
}

func (s UserStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.UserRepository.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": listOfUserIds}}); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting users (customers): %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
