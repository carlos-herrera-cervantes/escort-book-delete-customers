package db

import (
	"context"
	"log"
	"sync"
	"time"

	"escort-book-delete-customers/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *MongoClient
var singleMongoClient sync.Once

type MongoClient struct {
	PamentDB     *mongo.Database
	AuthorizerDB *mongo.Database
	SchedulerDB  *mongo.Database
}

func initMongoClient() {
	uri := config.InitializeMongo().Host
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctxWithTimeout, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("Error initializing mongo client: %s", err.Error())
	}

	mongoInstance = &MongoClient{
		PamentDB:     client.Database(config.InitializeMongo().Databases.Payment),
		AuthorizerDB: client.Database(config.InitializeMongo().Databases.Authorizer),
		SchedulerDB:  client.Database(config.InitializeMongo().Databases.Scheduler),
	}
}

func NewMongoClient() *MongoClient {
	singleMongoClient.Do(initMongoClient)
	return mongoInstance
}
