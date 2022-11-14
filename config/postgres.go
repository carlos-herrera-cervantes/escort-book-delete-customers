package config

import "os"

type postgres struct {
	Databases postgresDatabases
}

type postgresDatabases struct {
	CustomerProfile string
	EscortProfile   string
}

var singlePostgres *postgres

func InitializePostgres() *postgres {
	if singlePostgres != nil {
		return singlePostgres
	}

	lock.Lock()
	defer lock.Unlock()

	singlePostgres = &postgres{
		Databases: postgresDatabases{
			CustomerProfile: os.Getenv("ESCORT_BOOK_CUSTOMER_PROFILE_DB"),
			EscortProfile:   os.Getenv("ESCORT_BOOK_ESCORT_PROFILE_DB"),
		},
	}

	return singlePostgres
}
