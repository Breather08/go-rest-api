package database

var postTableSQL = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		author_id INTEGER NOT NULL,
		username TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		likes INTEGER NOT NULL,
		dislikes INTEGER NOT NULL
	);
`

var insertPostSQL = `
	INSERT INTO posts 
		(author_id, username, title, content, created_at, likes, dislikes)
		VALUES (?, ?, ?, ?, ?, 0, 0);
`
