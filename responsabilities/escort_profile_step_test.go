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

func TestEscortProfileStepExecute(t *testing.T) {
	controller := gomock.NewController(t)
	mockEscortProfileRepository := mockRepositories.NewMockIEscortProfileRepository(controller)
	mockCustomerRemovalStep := mocks.NewMockIAccountRemovalStep(controller)
	escortProfileStep := EscortProfileStep{
		EscortProfileRepository: mockEscortProfileRepository,
		Next:                    mockCustomerRemovalStep,
	}

	t.Run("Should log error when a repository error occurs", func(t *testing.T) {
		mockEscortProfileRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockCustomerRemovalStep.
			EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Times(1)

		escortProfileStep.Execute(context.Background(), []models.CustomerRemoval{
			{Id: primitive.NewObjectID()},
		})
	})
}
