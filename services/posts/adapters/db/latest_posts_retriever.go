package dbadapters

import (
	"database/sql"

	"example.com/services/posts/adapters/db/dto"
	"example.com/services/posts/core/domain"
	prepared_sql "example.com/services/posts/sql/prepared"
	"github.com/google/uuid"
)

type latestPostsRetriever struct {
	db *sql.DB
}

func NewLatestPostsRetriever(
	db *sql.DB,
) *latestPostsRetriever {
	return &latestPostsRetriever{
		db: db,
	}
}

func (s *latestPostsRetriever) GetLatestPosts(
	count int,
) ([]domain.Post, error) {
	var statement *sql.Stmt
	if preparedStatement, err := s.db.Prepare(
		prepared_sql.LatestPosts,
	); err != nil {
		return nil, err
	} else {
		statement = preparedStatement
	}
	defer statement.Close()

	var rows *sql.Rows
	if queryRows, err := statement.Query(count); err != nil {
		return []domain.Post{}, err
	} else {
		rows = queryRows
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post, err := dto.NewPostFromRows(rows)
		if err != nil {
			return []domain.Post{}, nil
		}

		postId, err := uuid.Parse(post.Id)
		if err != nil {
			return []domain.Post{}, nil
		}

		userId, err := uuid.Parse(post.Id)
		if err != nil {
			return []domain.Post{}, nil
		}

		posts = append(posts, domain.NewPost(
			postId,
			userId,
			post.Content,
			post.Timestamp,
		))
	}

	return posts, nil
}
