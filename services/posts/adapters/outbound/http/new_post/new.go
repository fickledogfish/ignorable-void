package newposthttpoutboundadapter

import "database/sql"

func New(
	db *sql.DB,
) *newPostInserter {
	return &newPostInserter{
		db: db,
	}
}
