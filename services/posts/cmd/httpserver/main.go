package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"example.com/core/environment"
	latestpostsdbinboundadapter "example.com/services/posts/adapters/inbound/db/latest_posts"
	insertposthttpinboundadapter "example.com/services/posts/adapters/inbound/http/insert_post"
	latestpostshttpinboundadapter "example.com/services/posts/adapters/inbound/http/latest_posts"
	newposthttpoutboundadapter "example.com/services/posts/adapters/outbound/http/new_post"
	insertpostserv "example.com/services/posts/core/services/insert_post"
	latestpostsserv "example.com/services/posts/core/services/latest_posts"
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

	log.Fatal(newHttpHandler(
		environment.ValueFor(envPort),
		insertposthttpinboundadapter.New(
			insertpostserv.New(
				newposthttpoutboundadapter.New(connection),
			),
		),
		latestpostshttpinboundadapter.New(
			latestpostsserv.New(
				latestpostsdbinboundadapter.New(connection),
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
