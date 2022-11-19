package responsabilities

import (
	"context"
	"fmt"

	"escort-book-delete-customers/models"
	"escort-book-delete-customers/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerRemovalStep struct {
	CustomerRemovalRepository repositories.ICustomerRemovalRepository
}

func (s CustomerRemovalStep) Execute(ctx context.Context, removals []models.CustomerRemoval) {
	listOfUserIds := convertUserIdToSlice(removals)

	if _, err := s.CustomerRemovalRepository.Update(
		ctx,
		bson.M{"user_id": bson.M{"$in": listOfUserIds}},
		map[string]bool{"executed": true},
	); err != nil {
		logger.Warn(fmt.Sprintf("Error updating customer removal: %s", err.Error()))
	}
}
