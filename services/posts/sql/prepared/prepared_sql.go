package prepared_sql

import _ "embed"

var (
	//go:embed insert_post.sql
	InsertPost string

	//go:embed select_latest_posts.sql
	LatestPosts string
)
