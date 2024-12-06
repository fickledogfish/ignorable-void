package ports

import (
	"example.com/core/uuid"
	"example.com/services/auth/core/domain"
)

type UserStorer interface {
	StoreUser(uuid.UUID, domain.SignUpRequest) error
}
