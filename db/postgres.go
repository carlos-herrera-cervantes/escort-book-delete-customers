package db

import (
	"database/sql"
	"log"
	"sync"

	"escort-book-delete-customers/config"

	_ "github.com/lib/pq"
)

var postgresInstance *PostgresClient
var singlePostgresClient sync.Once

type PostgresClient struct {
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

	postgresInstance = &PostgresClient{
		CustomerProfileDB: customerProfileDB,
		EscortProfileDB:   escortProfileDB,
	}
}

func NewPostgresClient() *PostgresClient {
	singlePostgresClient.Do(initPostgresClient)
	return postgresInstance
}
