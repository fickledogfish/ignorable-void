package ports

import (
	"example.com/services/posts/core/domain"
)

type NewPostInserter interface {
	InsertNewPost(domain.Post) error
}
