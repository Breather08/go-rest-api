package models

import "time"

// User model
type User struct {
	ID             int       `json:"id,string,omitempty"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"password"`
	CreatedAt        time.Time `json:"created_at"`
	Active         bool      `json:"active"`
	Rating         int       `json:"rating"`
}

// Comment model
type Comment struct {
	ID        int    `json:"id,string,omitempty"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	Author    User   `json:"author"`
	CreatedAt string `json:"created_at"`
}

// Post model
type Post struct {
	Author    User      `json:"author"`
	ID        int       `json:"id,string,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Active    bool      `json:"active"`
	Comments  []Comment `json:"comments"`
}
