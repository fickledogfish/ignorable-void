package ports

import (
	"example.com/core/uuid"
	"example.com/services/auth/core/domain"
)

type SignUpService interface {
	SignUp(
		credentials domain.SignUpRequest,
	) (uuid.UUID, error)
}
