package ports

import (
	"example.com/core/uuid"
	"example.com/services/posts/core/domain"
)

type InsertPostService interface {
	Post(domain.Post) (uuid.UUID, error)
}
