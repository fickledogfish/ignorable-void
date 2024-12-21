package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"example.com/core/environment"
	userretrieverdbinboundadapter "example.com/services/auth/adapters/inbound/db/user_retriever"
	signinhttpinboundadapter "example.com/services/auth/adapters/inbound/http/signin"
	signuphttpinboundadapter "example.com/services/auth/adapters/inbound/http/signup"
	userstorerdboutboundadapter "example.com/services/auth/adapters/outbound/db/user_storer"
	signinserv "example.com/services/auth/core/services/signin"
	signupserv "example.com/services/auth/core/services/signup"
)

const (
	envPort       = "PORT"
	envPGuserName = "POSTGRES_USER"
	envPGpassword = "POSTGRES_PASSWORD"
	envPGBDname   = "POSTGRES_DB"
)

func main() {
	dbConnection, err := connectToDatabase()
	if err != nil {
		panic(err)
	}

	log.Fatal(newHttpHandler(
		environment.ValueFor(envPort),
		signinhttpinboundadapter.New(
			signinserv.New(
				userretrieverdbinboundadapter.New(dbConnection),
			),
		),
		signuphttpinboundadapter.New(
			signupserv.New(
				userstorerdboutboundadapter.New(dbConnection),
				userretrieverdbinboundadapter.New(dbConnection),
			),
		),
	).ListenAndServe())
}

func connectToDatabase() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=postgres_db sslmode=disable",
		environment.ValueFor(envPGBDname),
		environment.ValueFor(envPGuserName),
		environment.ValueFor(envPGpassword),
	)

	return sql.Open("postgres", connectionString)
}
