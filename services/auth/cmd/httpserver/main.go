package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"example.com/core/environment"
	dbadapters "example.com/services/auth/adapters/db"
	httpadapters "example.com/services/auth/adapters/http"
	signinserv "example.com/services/auth/core/services/signin"
	signupserv "example.com/services/auth/core/services/signup"
	httphandler "example.com/services/auth/handlers/http"
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

	log.Fatal(httphandler.NewHttpHandler(
		environment.ValueFor(envPort),
		httpadapters.NewSignInHttpHandler(
			signinserv.New(
				dbadapters.NewUserRetriever(dbConnection),
			),
		),
		httpadapters.NewSignUpHttpHandler(
			signupserv.New(
				dbadapters.NewUserStorer(dbConnection),
				dbadapters.NewUserRetriever(dbConnection),
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
