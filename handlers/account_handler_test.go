package handlers

import (
	"context"
	"testing"

	"escort-book-delete-customers/strategies/mocks"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
)

func TestAccountHandlerHandleEvent(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAccountStrategyManager := mocks.NewMockIAccountStrategyManager(mockController)
	mockAccountStrategy := mocks.NewMockIAccountStrategy(mockController)
	accountHandler := AccountHandler{
		StrategyManager: mockAccountStrategyManager,
	}

	mockAccountStrategyManager.
		EXPECT().
		GetStrategy(gomock.Any()).
		Return(mockAccountStrategy).
		Times(1)
	mockAccountStrategy.
		EXPECT().
		SwitchAccountRemoval(gomock.Any(), gomock.Any()).
		Times(1)

	topic := "test-delete-customer"
	accountHandler.HandleEvent(context.Background(), &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
	})
}
