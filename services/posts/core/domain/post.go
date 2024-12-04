package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id uuid.UUID

	UserId    uuid.UUID
	Content   string
	Timestamp time.Time
}

func NewFreshPost(
	content string,
) Post {
	return Post{
		Content: content,
	}
}

func NewPost(
	id uuid.UUID,
	userId uuid.UUID,
	content string,
	timestamp time.Time,
) Post {
	return Post{
		Id:        id,
		UserId:    userId,
		Content:   content,
		Timestamp: timestamp,
	}
}
