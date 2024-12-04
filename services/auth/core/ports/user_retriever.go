package ports

import "example.com/services/auth/core/domain"

type UserRetriever interface {
	RetrieveUserWithName(string) (*domain.User, error)
}
