package strategies

import (
	"context"
	"encoding/json"
	"sync"

	"escort-book-delete-customers/repositories"
	"escort-book-delete-customers/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UndoRemovalStrategy struct {
	CustomerRemovalRepository repositories.ICustomerRemovalRepository
}

func (s UndoRemovalStrategy) SwitchAccountRemoval(ctx context.Context, value []byte) {
	var userActiveAccountEvent types.UserActiveAccountEvent

	if err := json.Unmarshal(value, &userActiveAccountEvent); err != nil {
		return
	}

	userId, _ := primitive.ObjectIDFromHex(userActiveAccountEvent.UserId)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		if err := s.CustomerRemovalRepository.Delete(ctx, bson.M{"user_id": userId}); err != nil {
			logger.Error("Error deleting a removal: ", err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
