package dto

import (
	"database/sql"
	"time"
)

type Post struct {
	Id string

	UserId    string
	Content   string
	Timestamp time.Time
}

func NewPostFromRows(
	row *sql.Rows,
) (post Post, err error) {
	err = row.Scan(
		&post.Id,
		&post.UserId,
		&post.Content,
		&post.Timestamp,
	)

	return
}
