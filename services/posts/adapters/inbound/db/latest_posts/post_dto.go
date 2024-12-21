package latestpostsdbinboundadapter

import (
	"database/sql"
	"time"
)

type postDto struct {
	Id string

	UserId    string
	Content   string
	Timestamp time.Time
}

func newPostFromRows(
	row *sql.Rows,
) (post postDto, err error) {
	err = row.Scan(
		&post.Id,
		&post.UserId,
		&post.Content,
		&post.Timestamp,
	)

	return
}
