SELECT
    id,
    user_id,
    content,
    creation_time
FROM posts ORDER BY creation_time DESC LIMIT $1;
