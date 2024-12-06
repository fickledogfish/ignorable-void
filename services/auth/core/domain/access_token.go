package domain

import "example.com/core/uuid"

type AccessToken struct {
	Id          uuid.UUID
	AccessToken string
}
