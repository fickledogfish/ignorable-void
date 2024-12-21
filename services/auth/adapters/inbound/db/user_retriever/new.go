package userretrieverdbinboundadapter

import "database/sql"

func New(
	db *sql.DB,
) *userRetriever {
	return &userRetriever{
		db: db,
	}
}
