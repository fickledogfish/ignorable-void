package userstorerdboutboundadapter

import "database/sql"

func New(
	db *sql.DB,
) *userStorer {
	return &userStorer{
		db: db,
	}
}
