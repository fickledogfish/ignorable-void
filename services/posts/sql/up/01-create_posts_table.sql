CREATE TABLE IF NOT EXISTS posts (
	id UUID PRIMARY KEY,

	user_id UUID,
	content VARCHAR(150),

	creation_time TIMESTAMP WITH TIME ZONE NOT NULL,

	FOREIGN KEY (user_id) REFERENCES users (id)
);
