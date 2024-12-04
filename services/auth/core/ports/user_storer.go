package ports

import (
	"example.com/services/auth/core/domain"

	"github.com/google/uuid"
)

type UserStorer interface {
	StoreUser(uuid.UUID, domain.SignUpRequest) error
}
