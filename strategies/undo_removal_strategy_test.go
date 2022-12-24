package strategies

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	mockRepositories "escort-book-delete-customers/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUndoRemovalStrategySwitchAccountRemoval(t *testing.T) {
	controller := gomock.NewController(t)
	mockCustomerRemovalRepository := mockRepositories.NewMockICustomerRemovalRepository(controller)
	undoRemovalStrategy := UndoRemovalStrategy{
		CustomerRemovalRepository: mockCustomerRemovalRepository,
	}

	t.Run("Should interrupt the process when unmarshal fails", func(t *testing.T) {
		value, err := json.Marshal("error")
		assert.NoError(t, err)

		undoRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})

	t.Run("Should log error when a repository error occurs", func(t *testing.T) {
		value := []byte(`{"userId": "638065d8fb0ce6eca0bf674c"}`)

		mockCustomerRemovalRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		undoRemovalStrategy.SwitchAccountRemoval(context.Background(), value)
	})
}
