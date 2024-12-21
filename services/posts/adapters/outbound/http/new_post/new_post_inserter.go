package newposthttpoutboundadapter

import (
	"database/sql"

	"example.com/services/posts/core/domain"
	prepared_sql "example.com/services/posts/sql/prepared"
)

type newPostInserter struct {
	db *sql.DB
}

func (s *newPostInserter) InsertNewPost(
	post domain.Post,
) error {
	var statement *sql.Stmt
	if preparedStatement, err := s.db.Prepare(
		prepared_sql.InsertPost,
	); err != nil {
		return err
	} else {
		statement = preparedStatement
	}
	defer statement.Close()

	if _, err := statement.Exec(
		post.Id,
		post.UserId,
		post.Content,
	); err != nil {
		return err
	}

	return nil
}
