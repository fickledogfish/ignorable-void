package ports

import (
	"example.com/services/auth/core/domain"
	"github.com/google/uuid"
)

type SignUpService interface {
	SignUp(
		credentials domain.SignUpRequest,
	) (uuid.UUID, error)
}
