package dto

import "database/sql"

type User struct {
	Id string

	Username string
	Password string
}

func NewUserFromRow(
	row *sql.Row,
) (user User, err error) {
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)

	return
}
