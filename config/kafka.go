package config

import (
	"os"
	"sync"
)

type kafka struct {
	BootstrapServers string
	GroupId          string
	Topics           topic
}

type topic struct {
	UserDeleteAccount string
	UserActiveAccount string
}

var lock = &sync.Mutex{}
var singleKafka *kafka

func InitializeKafka() *kafka {
	if singleKafka != nil {
		return singleKafka
	}

	lock.Lock()
	defer lock.Unlock()

	singleKafka = &kafka{
		BootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		GroupId:          os.Getenv("KAFKA_GROUP_ID"),
		Topics: topic{
			UserDeleteAccount: "user-delete-account",
			UserActiveAccount: "user-active-account",
		},
	}

	return singleKafka
}
