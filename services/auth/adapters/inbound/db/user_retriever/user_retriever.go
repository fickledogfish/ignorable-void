package userretrieverdbinboundadapter

import (
	"database/sql"

	"example.com/core/uuid"
	"example.com/services/auth/core/domain"
	prepared_sql "example.com/services/auth/sql/prepared"
)

type userRetriever struct {
	db *sql.DB
}

func (s *userRetriever) RetrieveUserWithName(
	username string,
) (*domain.User, error) {
	var statement *sql.Stmt
	if preparedStatement, err := s.db.Prepare(
		prepared_sql.UserWithName,
	); err != nil {
		return nil, err
	} else {
		statement = preparedStatement
	}
	defer statement.Close()

	var user userDto
	if storedUser, err := newUserFromRow(
		statement.QueryRow(username),
	); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}

		return nil, err
	} else {
		user = storedUser
	}

	var id uuid.UUID
	if userId, err := uuid.Parse(user.Id); err != nil {
		return nil, err
	} else {
		id = userId
	}

	newUser := domain.NewUser(
		id,
		user.Username,
		user.Password,
	)
	return &newUser, nil
}
