package handlers

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//go:generate mockgen -destination=./mocks/ievent_handler.go -package=mocks --build_flags=--mod=mod . IEventHandler
type IEventHandler interface {
	HandleEvent(ctx context.Context, message *kafka.Message)
}
