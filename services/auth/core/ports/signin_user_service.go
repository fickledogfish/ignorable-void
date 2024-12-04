package ports

import "example.com/services/auth/core/domain"

type SignInUserService interface {
	WithCredentials(
		credentials domain.UserSignInCredentials,
	) (domain.AccessToken, error)
}
