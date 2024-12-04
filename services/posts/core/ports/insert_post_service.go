package ports

import (
	"example.com/services/posts/core/domain"
	"github.com/google/uuid"
)

type InsertPostService interface {
	Post(domain.Post) (uuid.UUID, error)
}
