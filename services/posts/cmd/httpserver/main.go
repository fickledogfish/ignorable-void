package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"example.com/core/environment"
	dbadapters "example.com/services/posts/adapters/db"
	httpadapters "example.com/services/posts/adapters/http"
	insertpostserv "example.com/services/posts/core/services/insert_post"
	latestpostsserv "example.com/services/posts/core/services/latest_posts"
	httphandler "example.com/services/posts/handlers/http"
)

const (
	envPort       = "PORT"
	envPGuserName = "POSTGRES_USER"
	envPGpassword = "POSTGRES_PASSWORD"
	envPGBDname   = "POSTGRES_DB"
)

func main() {
	connection, err := connectToDatabase()
	if err != nil {
		panic(err)
	}

	log.Fatal(httphandler.NewHttpHandler(
		environment.ValueFor(envPort),
		httpadapters.NewInsertPostHttpHandler(
			insertpostserv.New(
				dbadapters.NewPostInserter(connection),
			),
		),
		httpadapters.NewLatestPostsHttpHandler(
			latestpostsserv.New(
				dbadapters.NewLatestPostsRetriever(connection),
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
