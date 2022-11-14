package main

import (
	"fmt"

	"escort-book-delete-customers/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
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
			fmt.Println("New message")
		case kafka.PartitionEOF:
			fmt.Println("Reached: ", e)
		case kafka.Error:
			fmt.Println("Error: ", e)
			run = false
		}
	}

	_ = consumer.Close()
}
