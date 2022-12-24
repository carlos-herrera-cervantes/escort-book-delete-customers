package strategies

import (
	"context"
	"errors"
	"testing"

	"escort-book-delete-customers/models"
	mockRepositories "escort-book-delete-customers/repositories/mocks"

	"github.com/golang/mock/gomock"
)

func TestDoRemovalStrategySwitchAccountRemoval(t *testing.T) {
	controller := gomock.NewController(t)
	mockCustomerRemovalRepository := mockRepositories.NewMockICustomerRemovalRepository(controller)
	mockCustomerProfileRepository := mockRepositories.NewMockICustomerProfileRepository(controller)
	mockEscortProfileRepository := mockRepositories.NewMockIEscortProfileRepository(controller)
	doRemovalStrategy := DoRemovalStrategy{
		CustomerRemovalRepository: mockCustomerRemovalRepository,
		CustomerProfileRepository: mockCustomerProfileRepository,
		EscortProfileRepository:   mockEscortProfileRepository,
	}

	t.Run("Should interrupt the process when unmarshal fails", func(t *testing.T) {
		value := []byte("error")
		doRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})

	t.Run("Should interrupt the process when customer does not exists", func(t *testing.T) {
		value := []byte(`{"userId": "63806dce5b95019383804f1e", "userType": "Customer", "userEmail": "user@example.com"}`)

		mockCustomerProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CustomerProfile{}, errors.New("dummy error")).
			Times(1)

		doRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})

	t.Run("Should interrupt the process when escort does not exists", func(t *testing.T) {
		value := []byte(`{"userId": "63806dce5b95019383804f1e", "userType": "Escort", "userEmail": "user@example.com"}`)

		mockEscortProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.EscortProfile{}, errors.New("dummy error")).
			Times(1)

		doRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})

	t.Run("Should log error when customer removal repository fails", func(t *testing.T) {
		value := []byte(`{"userId": "63806dce5b95019383804f1e", "userType": "Escort", "userEmail": "user@example.com"}`)

		mockEscortProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.EscortProfile{}, nil).
			Times(1)
		mockCustomerRemovalRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.CustomerRemoval{}, errors.New("dummy error")).
			Times(1)

		doRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})
}
