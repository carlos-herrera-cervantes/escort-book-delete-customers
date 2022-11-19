package handlers

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IEventHandler interface {
	HandleEvent(ctx context.Context, message *kafka.Message)
}
