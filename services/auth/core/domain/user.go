package domain

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID

	Username string
	Password string
}

func NewUser(
	id uuid.UUID,
	username string,
	password string,
) User {
	return User{
		Id: id,

		Username: username,
		Password: password,
	}
}
