package prepared_sql

import _ "embed"

var (
	//go:embed insert_user.sql
	InsertUser string

	//go:embed select_user_with_name.sql
	UserWithName string
)
