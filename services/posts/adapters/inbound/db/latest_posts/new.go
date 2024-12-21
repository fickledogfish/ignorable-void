package latestpostsdbinboundadapter

import "database/sql"

func New(
	db *sql.DB,
) *latestPostsRetriever {
	return &latestPostsRetriever{
		db: db,
	}
}
