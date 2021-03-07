package database

import (
	"encoding/json"
	"fmt"
	"forum/backend/server/models"
	"log"
	"net/http"
)

func CreatePostsTable() {
	postsTable, err := db.Prepare(postTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = postsTable.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Post table created in DB")
}

func insertPostIntoDB(postData models.Post) {
	insertPost, err := db.Prepare(insertPostSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertPost.Exec(
		postData.Author.ID,
		postData.Author.Username,
		postData.Title,
		postData.Content,
		postData.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New post inserted into DB")
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	insertPostIntoDB(post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New post sent to clientside")
}

func getAllPosts() []models.Post {
	var posts []models.Post

	row, err := db.Query(`
		SELECT * FROM posts
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var author models.User

	for row.Next() {
		var post = models.Post{
			Author: author,
		}
		err = row.Scan(
			&post.ID,
			&post.Author.ID,
			&post.Author.Username,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
		)

		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	return posts
}

func SendAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	posts := getAllPosts()
	json.NewEncoder(w).Encode(posts)

	fmt.Println("All posts sent to clientside")
}
