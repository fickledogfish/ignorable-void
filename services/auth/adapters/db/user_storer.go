package dbadapters

import (
	"database/sql"

	"github.com/google/uuid"

	"example.com/services/auth/core/domain"
	prepared_sql "example.com/services/auth/sql/prepared"
)

type userStorer struct {
	db *sql.DB
}

func NewUserStorer(
	db *sql.DB,
) *userStorer {
	return &userStorer{
		db: db,
	}
}

func (s *userStorer) StoreUser(
	id uuid.UUID,
	newUser domain.SignUpRequest,
) error {
	var statement *sql.Stmt
	if preparedStatement, err := s.db.Prepare(
		prepared_sql.InsertUser,
	); err != nil {
		return err
	} else {
		statement = preparedStatement
	}
	defer statement.Close()

	if _, err := statement.Exec(
		id,
		newUser.Username,
		newUser.Password,
	); err != nil {
		return err
	}

	return nil
}
