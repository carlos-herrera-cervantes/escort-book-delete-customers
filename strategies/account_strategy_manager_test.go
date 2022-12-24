package strategies

import (
	"testing"

	mockRepositories "escort-book-delete-customers/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAccountStrategyManagerGetStrategy(t *testing.T) {
	controller := gomock.NewController(t)
	mockCustomerRemovalRepository := mockRepositories.NewMockICustomerRemovalRepository(controller)
	undoRemovalStrategy := UndoRemovalStrategy{
		CustomerRemovalRepository: mockCustomerRemovalRepository,
	}
	strategies := map[string]IAccountStrategy{
		"dummy-topic": undoRemovalStrategy,
	}
	accountStrategyManager := AccountStrategyManager{
		Strategies: strategies,
	}

	t.Run("Should return the correct strategy", func(t *testing.T) {
		strategy := accountStrategyManager.GetStrategy("dummy-topic")
		assert.NotNil(t, strategy)
	})
}
