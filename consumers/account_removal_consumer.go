package consumers

import (
	"context"

	"escort-book-delete-customers/config"
	"escort-book-delete-customers/handlers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/inconshreveable/log15"
)

type AccountRemovalConsumer struct {
	EventHandler handlers.IEventHandler
}

var logger = log.New("consumers")

func (c AccountRemovalConsumer) StartConsumer() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.InitializeKafka().BootstrapServers,
		"group.id":           config.InitializeKafka().GroupId,
		"auto.offset.reset":  "smallest",
		"enable.auto.commit": true,
	})
	topics := []string{
		config.InitializeKafka().Topics.UserActiveAccount,
		config.InitializeKafka().Topics.UserDeleteAccount,
	}
	_ = consumer.SubscribeTopics(topics, nil)
	run := true

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			c.EventHandler.HandleEvent(context.Background(), e)
			logger.Info("New message arrived")
		case kafka.PartitionEOF:
			logger.Info("Reached: ", e)
		case kafka.Error:
			logger.Error("Kafka error: ", e.Error())
			run = false
		}
	}

	_ = consumer.Close()
}
