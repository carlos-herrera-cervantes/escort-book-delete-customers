package db

import (
	"database/sql"
	"escort-book-delete-customers/config"
	"log"
	"sync"
)

var postgresInstance *postgresClient
var singlePostgresClient sync.Once

type postgresClient struct {
	CustomerProfileDB *sql.DB
	EscortProfileDB   *sql.DB
}

func initPostgresClient() {
	customerProfileDB, err := sql.Open("postgres", config.InitializePostgres().Databases.CustomerProfile)

	if err != nil {
		log.Fatalf("Error connectiing with customer_profile DB: %s", err.Error())
	}

	escortProfileDB, err := sql.Open("postgres", config.InitializePostgres().Databases.EscortProfile)

	if err != nil {
		log.Fatalf("Error connectiing with escort_profile DB: %s", err.Error())
	}

	postgresInstance = &postgresClient{
		CustomerProfileDB: customerProfileDB,
		EscortProfileDB:   escortProfileDB,
	}
}

func NewPostgresClient() *postgresClient {
	singlePostgresClient.Do(initPostgresClient)
	return postgresInstance
}
