package database

const postTableSQL = `
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

// password --> hashed_password char(60)
const userTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		hashed_password CHAR(60) NOT NULL,
		created_at DATE NOT NULL,
		rating TEXT NOT NULL,
		active BOOLEAN DEFAULT TRUE
	);
`

const insertPostSQL = `
	INSERT INTO posts 
		(author_id, username, title, content, created_at, likes, dislikes)
		VALUES (?, ?, ?, ?, ?, 0, 0);
`

const insertUserSQL = `
	INSERT INTO users 
		(username, email, hashed_password, created_at, active, rating)
		VALUES (?, ?, ?, ?, ?, 0);
`
