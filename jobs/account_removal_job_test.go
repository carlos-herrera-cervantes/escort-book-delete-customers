package jobs

import (
	"context"
	"errors"
	"testing"

	"escort-book-delete-customers/models"
	mockRepositories "escort-book-delete-customers/repositories/mocks"

	"github.com/golang/mock/gomock"
)

func TestAccountRemovalJobRemoveEscorts(t *testing.T) {
	controller := gomock.NewController(t)

	mockCustomerRemovalRepository := mockRepositories.NewMockICustomerRemovalRepository(controller)
	mockCustomerProfileRepository := mockRepositories.NewMockICustomerProfileRepository(controller)
	mockEscortProfileRepository := mockRepositories.NewMockIEscortProfileRepository(controller)
	mockCardRepository := mockRepositories.NewMockICardRepository(controller)
	mockAccessTokenRepository := mockRepositories.NewMockIAccessTokenRepository(controller)
	mockBankAccountRepository := mockRepositories.NewMockIBankAccountRepository(controller)
	mockUserPaymentRepository := mockRepositories.NewMockIUserPaymentRepository(controller)
	mockUserRepository := mockRepositories.NewMockIUserRepository(controller)

	accountRemovalJob := AccountRemovalJob{
		CustomerRemovalRepository: mockCustomerRemovalRepository,
		CustomerProfileRepository: mockCustomerProfileRepository,
		EscortProfileRepository:   mockEscortProfileRepository,
		CardRepository:            mockCardRepository,
		AccessTokenRepository:     mockAccessTokenRepository,
		BankAccountRepository:     mockBankAccountRepository,
		UserPaymentRepository:     mockUserPaymentRepository,
		UserRepository:            mockUserRepository,
	}

	t.Run("Should interrupt the process when customer removal repository fails", func(t *testing.T) {
		mockCustomerRemovalRepository.
			EXPECT().
			GetMany(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CustomerRemoval{}, errors.New("dummy error")).
			Times(1)

		accountRemovalJob.removeEscorts(context.Background())
	})

	t.Run("Should interrupt the process when there are no removals", func(t *testing.T) {
		mockCustomerRemovalRepository.
			EXPECT().
			GetMany(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CustomerRemoval{}, nil).
			Times(1)

		accountRemovalJob.removeEscorts(context.Background())
	})
}
