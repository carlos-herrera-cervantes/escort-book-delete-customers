package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardStep struct {
	CardRepository repositories.ICardRepository
	Next           IAccountRemovalStep
}

func (s CardStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if err := s.CardRepository.DeleteMany(ctx, bson.M{"customerId": bson.M{"$in": listOfUserIds}}); err != nil {
		logger.Warn(fmt.Sprintf("Error deleting customer cards: %s", err.Error()))
	}

	s.Next.Execute(ctx, removals)
}

func convertUserIdToSlice(elements []models.CustomerRemoval) []primitive.ObjectID {
	userIds := []primitive.ObjectID{}

	for _, value := range elements {
		userIds = append(userIds, value.UserId)
	}

	return userIds
}
