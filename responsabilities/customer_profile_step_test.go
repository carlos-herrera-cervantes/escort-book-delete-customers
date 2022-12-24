package responsabilities

import (
	"context"
	"errors"
	"testing"

	"escort-book-delete-customers/models"
	mockRepositories "escort-book-delete-customers/repositories/mocks"
	"escort-book-delete-customers/responsabilities/mocks"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerProfileStepExecute(t *testing.T) {
	controller := gomock.NewController(t)
	mockCustomerProfileRepository := mockRepositories.NewMockICustomerProfileRepository(controller)
	mockCustomerRemovalStep := mocks.NewMockIAccountRemovalStep(controller)
	customerProfileStep := CustomerProfileStep{
		CustomerProfileRepository: mockCustomerProfileRepository,
		Next:                      mockCustomerRemovalStep,
	}

	t.Run("Should log error when a repository error occurs", func(t *testing.T) {
		mockCustomerProfileRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockCustomerRemovalStep.
			EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Times(1)

		customerProfileStep.Execute(context.Background(), []models.CustomerRemoval{
			{Id: primitive.NewObjectID()},
		})
	})
}
