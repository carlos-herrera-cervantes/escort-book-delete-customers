package handlers

import (
	"context"

	"escort-book-delete-customers/strategies"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AccountHandler struct {
	StrategyManager strategies.IAccountStrategyManager
}

func (h AccountHandler) HandleEvent(ctx context.Context, message *kafka.Message) {
	topic := message.TopicPartition.Topic
	strategy := h.StrategyManager.GetStrategy(*topic)
	strategy.SwitchAccountRemoval(ctx, message.Value)
}
