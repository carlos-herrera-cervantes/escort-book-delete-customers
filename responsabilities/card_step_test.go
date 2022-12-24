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

func TestCardStepExecute(t *testing.T) {
	controller := gomock.NewController(t)
	mockCardRepository := mockRepositories.NewMockICardRepository(controller)
	mockCustomerRemovalStep := mocks.NewMockIAccountRemovalStep(controller)
	cardStep := CardStep{
		CardRepository: mockCardRepository,
		Next:           mockCustomerRemovalStep,
	}

	t.Run("Should log error when a repository error occurs", func(t *testing.T) {
		mockCardRepository.
			EXPECT().
			DeleteMany(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockCustomerRemovalStep.
			EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Times(1)

		cardStep.Execute(context.Background(), []models.CustomerRemoval{
			{Id: primitive.NewObjectID()},
		})
	})
}
