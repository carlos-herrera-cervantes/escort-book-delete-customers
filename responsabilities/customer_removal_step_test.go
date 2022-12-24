package responsabilities

import (
	"context"
	"errors"
	"testing"

	"escort-book-delete-customers/models"
	mockRepositories "escort-book-delete-customers/repositories/mocks"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerRemovalStepExecute(t *testing.T) {
	controller := gomock.NewController(t)
	mockCustomerRemovalRepository := mockRepositories.NewMockICustomerRemovalRepository(controller)
	customerRemovalStep := CustomerRemovalStep{
		CustomerRemovalRepository: mockCustomerRemovalRepository,
	}

	t.Run("Should log error when a repository error occurs", func(t *testing.T) {
		mockCustomerRemovalRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.CustomerRemoval{}, errors.New("dummy error")).
			Times(1)

		customerRemovalStep.Execute(context.Background(), []models.CustomerRemoval{
			{Id: primitive.NewObjectID()},
		})
	})
}
