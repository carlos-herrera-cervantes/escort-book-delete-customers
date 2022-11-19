package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type AccessTokenStep struct {
	AccessTokenRepository repositories.IAccessTokenRepository
	Next                  IAccountRemovalStep
}

func (s AccessTokenStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.AccessTokenRepository.DeleteMany(ctx, bson.M{"userId": bson.M{"$in": listOfUserIds}}); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting customer tokens: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}
