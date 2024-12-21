package userretrieverdbinboundadapter

import "database/sql"

type userDto struct {
	Id string

	Username string
	Password string
}

func newUserFromRow(
	row *sql.Row,
) (user userDto, err error) {
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)

	return
}
