package router

import (
	"forum/backend/server/database"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/posts", database.SendAllPosts)
	mux.HandleFunc("/api/create-post", database.InsertPost)

	return mux
}
