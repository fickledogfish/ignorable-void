package domain

import "github.com/google/uuid"

type AccessToken struct {
	Id          uuid.UUID
	AccessToken string
}
